package apiserver

import (
	"net/http"

	"github.com/abd-rakhman/qysqa-back/internal/db/sqlc"
)

func (s *server) getLastContestsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	contest, err := s.database.GetLastTwoContests(ctx)
	if err != nil {
		http.Error(w, "Failed to create reading: "+err.Error(), http.StatusInternalServerError)
	}

	s.respond(w, r, http.StatusAccepted, ParseContests(contest))
}

type ResponseReading struct {
	Id   int64  `json:"id"`
	Text string `json:"text"`
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
	Id      int64           `json:"id"`
	Reading ResponseReading `json:"reading"`
	Diktant DiktantResponse `json:"diktant"`
	Speech  SpeechResponse  `json:"speech"`
}

func ParseContests(contests []sqlc.GetLastTwoContestsRow) []ContestResponse {
	var response []ContestResponse
	for _, contest := range contests {
		response = append(response, ContestResponse{
			Id: contest.ID,
			Reading: ResponseReading{
				Id:   contest.ReadingID,
				Text: contest.ReadingText.String, // Assuming the text is valid and not null
			},
			Diktant: DiktantResponse{
				Id:   contest.DiktantID,
				Text: contest.DiktantText.String, // Assuming the text is valid and not null
			},
			Speech: SpeechResponse{
				Id:   contest.SpeechesID,
				Text: contest.SpeechesText.String, // Assuming the text is valid and not null
			},
		})
	}
	return response
}
