package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/abd-rakhman/qysqa-back/internal/db/sqlc"
)

func (s *server) getLastContestsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	contest, err := s.database.GetLastTwoContests(ctx)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	responseBody, err := ParseContests(contest)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, r, http.StatusAccepted, responseBody)
}

type VariantResponse struct {
	Option string `json:"option"`
}

type QuestionResponse struct {
	Text     string            `json:"text"`
	Question string            `json:"question"`
	Variants []VariantResponse `json:"variants"`
}

type ResponseReading struct {
	Id        int64              `json:"id"`
	Text      string             `json:"text"`
	Questions []QuestionResponse `json:"questions"`
}

type DiktantResponse struct {
	Id   int64  `json:"id"`
	Text string `json:"text"`
}

type SpeechResponse struct {
	Id   int64  `json:"id"`
	Text string `json:"text"`
}

type ContestResponse struct {
	Id        int64           `json:"id"`
	Reading   ResponseReading `json:"reading"`
	Diktant   DiktantResponse `json:"diktant"`
	Speech    SpeechResponse  `json:"speech"`
	StartTime int64           `json:"start_time"`
	EndTime   int64           `json:"end_time"`
}

func ParseContests(contests []sqlc.GetLastTwoContestsRow) ([]ContestResponse, error) {
	var response []ContestResponse
	for _, contest := range contests {
		var questions []QuestionResponse
		err := json.Unmarshal(contest.Questions, &questions)
		if err != nil {
			return nil, err
		}

		response = append(response, ContestResponse{
			Id: contest.ContestID,
			Reading: ResponseReading{
				Id:        contest.ReadingID,
				Text:      contest.ReadingText.String, // Assuming the text is valid and not null
				Questions: questions,
			},
			Diktant: DiktantResponse{
				Id:   contest.DiktantID,
				Text: contest.DiktantText.String, // Assuming the text is valid and not null
			},
			Speech: SpeechResponse{
				Id:   contest.SpeechID,
				Text: contest.SpeechText.String, // Assuming the text is valid and not null
			},
			StartTime: contest.StartTime.Time.UnixMilli(),
			EndTime:   contest.EndTime.Time.UnixMilli(),
		})
	}
	return response, nil
}
