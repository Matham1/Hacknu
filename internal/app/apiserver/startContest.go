package apiserver

// import (
// 	"encoding/json"
// 	"errors"
// 	"net/http"
// )

// type StartContestRequest struct {
// 	ContestId int64 `json:"contest_id"`
// 	UserId    int64 `json:"user_id"`
// }

// func (s *server) startContest(w http.ResponseWriter, r *http.Request) {
// 	var req StartContestRequest
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&req); err != nil {
// 		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	if err := validateStartContestRequest(req); err != nil {
// 		http.Error(w, "Validation error: "+err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	// contest, err := s.database.StartContest(ctx, req)
// 	// if err != nil {
// 	// 	s.error(w, r, http.StatusInternalServerError, err)
// 	// 	return
// 	// }

// 	responseBody, err := ParseContest(contest)
// 	if err != nil {
// 		s.error(w, r, http.StatusInternalServerError, err)
// 		return
// 	}

// 	s.respond(w, r, http.StatusAccepted, responseBody)
// }

// func validateStartContestRequest(req StartContestRequest) error {
// 	if req.ContestId == 0 {
// 		return errors.New("contest_id must be greater than 0")
// 	}

// 	if req.UserId == 0 {
// 		return errors.New("user_id must be greater than 0")
// 	}

// 	return nil
// }
