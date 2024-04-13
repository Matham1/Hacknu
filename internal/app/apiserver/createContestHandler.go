package apiserver

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/abd-rakhman/qysqa-back/internal/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

// I know, MongoDB is more appropriate for such relations, but I don't have time to learn it(
type Variant struct {
	Option      string `json:"option"`
	IsCorrect   bool   `json:"is_correct"`
	Explanation string `json:"explanation"`
}

type Question struct {
	Text     string    `json:"text"`
	Question string    `json:"question"`
	Variants []Variant `json:"variants"`
}

type Reading struct {
	Text      string     `json:"text"`
	Questions []Question `json:"questions"`
}

type CreateContestRequest struct {
	Reading  Reading `json:"reading"`
	Diktant  string  `json:"diktant"`
	Speeches string  `json:"speeches"`
	StartAt  string  `json:"start_at"`
	EndAt    string  `json:"end_at"`
}

type Contest struct {
	ReadingID  int64     `json:"reading_id"`
	DiktantID  int64     `json:"diktant_id"`
	SpeechesID int64     `json:"speeches_id"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
}

func (s *server) createContest(w http.ResponseWriter, r *http.Request) {
	var req CreateContestRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := validateCreateContestRequest(req); err != nil {
		http.Error(w, "Validation error: "+err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	// TODO: Add transaction
	reading, err := s.database.CreateReading(ctx, pgtype.Text{String: req.Reading.Text, Valid: true})
	if err != nil {
		http.Error(w, "Failed to create reading: "+err.Error(), http.StatusInternalServerError)
		return
	}

	for _, question := range req.Reading.Questions {
		questionResult, err := s.database.CreateReadingQuestion(ctx, sqlc.CreateReadingQuestionParams{
			Text:      question.Text,
			Question:  question.Question,
			ReadingID: pgtype.Int8{Int64: int64(reading.ID), Valid: true},
		})
		if err != nil {
			http.Error(w, "Failed to create reading question: "+err.Error(), http.StatusInternalServerError)
			return
		}

		for _, variant := range question.Variants {
			_, err := s.database.CreateQuestionVariant(ctx, sqlc.CreateQuestionVariantParams{
				QuestionID:  pgtype.Int8{Int64: int64(questionResult.ID), Valid: true},
				Option:      variant.Option,
				IsCorrect:   variant.IsCorrect,
				Explanation: variant.Explanation,
			})
			if err != nil {
				http.Error(w, "Failed to create question variant: "+err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	diktant, err := s.database.CreateDiktant(ctx, pgtype.Text{String: req.Diktant, Valid: true})
	if err != nil {
		http.Error(w, "Failed to create diktant: "+err.Error(), http.StatusInternalServerError)
		return
	}

	speech, err := s.database.CreateSpeech(ctx, pgtype.Text{String: req.Speeches, Valid: true})
	if err != nil {
		http.Error(w, "Failed to create speech: "+err.Error(), http.StatusInternalServerError)
		return
	}

	startTime, errStart := time.Parse(time.RFC3339, req.StartAt)
	if errStart != nil {
		http.Error(w, "Invalid start time format", http.StatusBadRequest)
		return
	}
	endTime, errEnd := time.Parse(time.RFC3339, req.EndAt)
	if errEnd != nil {
		http.Error(w, "Invalid end time format", http.StatusBadRequest)
		return
	}

	contest, err := s.database.CreateContest(ctx, sqlc.CreateContestParams{
		ReadingID:  pgtype.Int8{Int64: int64(reading.ID), Valid: true},
		DiktantID:  pgtype.Int8{Int64: int64(diktant.ID), Valid: true},
		SpeechesID: pgtype.Int8{Int64: int64(speech.ID), Valid: true},
		StartTime:  pgtype.Timestamp{Time: startTime, Valid: true},
		EndTime:    pgtype.Timestamp{Time: endTime, Valid: true},
	})
	if err != nil {
		http.Error(w, "Failed to create contest: "+err.Error(), http.StatusInternalServerError)
		return
	}

	s.respond(w, r, http.StatusCreated, Contest{
		ReadingID:  contest.ReadingID.Int64,
		DiktantID:  contest.DiktantID.Int64,
		SpeechesID: contest.SpeechesID.Int64,
		StartTime:  contest.StartTime.Time,
		EndTime:    contest.EndTime.Time,
	})
}

// validateCreateContestRequest checks all aspects of the incoming request for creating a contest.
func validateCreateContestRequest(req CreateContestRequest) error {
	// Validate reading text
	if strings.TrimSpace(req.Reading.Text) == "" {
		return errors.New("reading text cannot be empty")
	}

	// Validate that there is at least one question in reading
	if len(req.Reading.Questions) == 0 {
		return errors.New("at least one question is required in a reading")
	}

	// Iterate through each question
	for _, question := range req.Reading.Questions {
		if strings.TrimSpace(question.Text) == "" {
			return errors.New("question text cannot be empty")
		}
		if strings.TrimSpace(question.Question) == "" {
			return errors.New("question description cannot be empty")
		}

		// Check each variant within a question
		if len(question.Variants) == 0 {
			return errors.New("at least one variant is required per question")
		}
		for _, variant := range question.Variants {
			if strings.TrimSpace(variant.Option) == "" {
				return errors.New("option text cannot be empty")
			}
			if strings.TrimSpace(variant.Explanation) == "" {
				return errors.New("explanation cannot be empty")
			}
		}
	}

	// Validate diktant and speeches
	if strings.TrimSpace(req.Diktant) == "" {
		return errors.New("diktant text cannot be empty")
	}
	if strings.TrimSpace(req.Speeches) == "" {
		return errors.New("speeches text cannot be empty")
	}

	// Validate date strings
	if err := validateDate(req.StartAt); err != nil {
		return errors.New("invalid start date: " + err.Error())
	}
	if err := validateDate(req.EndAt); err != nil {
		return errors.New("invalid end date: " + err.Error())
	}

	// Validate start is before end
	startAt, _ := time.Parse(time.RFC3339, req.StartAt)
	endAt, _ := time.Parse(time.RFC3339, req.EndAt)
	if startAt.After(endAt) {
		return errors.New("start time must be before end time")
	}

	return nil
}

// validateDate checks if a date string is properly formatted according to RFC3339.
func validateDate(dateStr string) error {
	if dateStr == "" {
		return errors.New("date string cannot be empty")
	}
	_, err := time.Parse(time.RFC3339, dateStr)
	return err
}
