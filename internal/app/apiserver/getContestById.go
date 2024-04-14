package apiserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/abd-rakhman/qysqa-back/internal/db/sqlc"
	"github.com/gorilla/mux"
)

func (s *server) getContestById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		s.error(w, r, http.StatusBadRequest, err)
		return
	}

	contest, err := s.database.GetContest(ctx, int64(id))
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	constestResponse, err := ParseContest(contest)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, r, http.StatusAccepted, constestResponse)
}

func ParseContest(contest sqlc.GetContestRow) (*ContestResponse, error) {
	var questions []QuestionResponse
	err := json.Unmarshal(contest.Questions, &questions)
	if err != nil {
		return nil, err
	}

	response := ContestResponse{
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
	}
	return &response, nil
}
