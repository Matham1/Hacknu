package apiserver

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/abd-rakhman/qysqa-back/internal/db/sqlc"
	"github.com/jackc/pgx/v5"
)

// Define the request structure
type CheckRequest struct {
	ID           int `json:"id"`
	ChosenOption int `json:"chosenOption"`
}

// Define the response structure
type CheckResponse struct {
	IsCorrect bool `json:"isCorrect"`
}

// Handler function for /test/check endpoint
func CheckHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the request body into CheckRequest struct
	var req CheckRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// TODO: Implement logic to check correctness based on req.ID and req.ChosenOption
	// For now, let's assume the chosen option is correct if it's equal to the ID (just for demonstration)

	// Find the question by ID
	questionid := -1
	var qcorrect int

	for _, text := range quizData.Texts {
		for _, q := range text.Questions {
			if q.ID == req.ID {
				questionid = q.ID
				qcorrect = q.Correct
				break
			}
		}
		if questionid != -1 {
			break
		}
	}

	// If question not found, return 404 Not Found
	if questionid == -1 {
		http.Error(w, "Question Not Found", http.StatusNotFound)
		return
	}

	// Check if the chosen option is correct
	isCorrect := req.ChosenOption == qcorrect

	// Prepare the response
	res := CheckResponse{IsCorrect: isCorrect}

	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Encode the response struct into JSON and write it to the response writer
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func Start(config AppConfig) error {
	db, closeDB, err := connectToDatabase(config.DATABASE_URL)
	if err != nil {
		return err
	}
	defer closeDB()

	srv := newServer(db)
	log.Printf("Server is listening on port %s...\n", config.BIND_ADDRESS)
	return http.ListenAndServe(config.BIND_ADDRESS, srv)
}

func connectToDatabase(databaseUrl string) (*sqlc.Queries, func(), error) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, databaseUrl)
	if err != nil {
		return nil, nil, err
	}
	close := func() {
		conn.Close(ctx)
	}

	return sqlc.New(conn), close, nil
}
