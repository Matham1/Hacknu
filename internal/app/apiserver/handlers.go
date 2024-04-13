package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func (s *server) submitCodeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}

		code := r.PostFormValue("code")
		stdin := r.PostFormValue("stdIn")
		userId, err := strconv.Atoi(r.PostFormValue("userId"))
		if err != nil {
			http.Error(w, "Invalid UserID", http.StatusBadRequest)
			return
		}

		if code == "" || userId == 0 {
			http.Error(w, "Code and UserID are required", http.StatusBadRequest)
			return
		}

		submission := Submission{
			UserId: userId,
			Code:   code,
			Stdin:  stdin,
		}

		if err := s.publishMessageQueue(submission); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, nil)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func (s *server) outputHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}
	uuid := r.PostFormValue("uuid")
	stdout := r.PostFormValue("stdout")
	sessionId := r.PostFormValue("sessionId")

	s.logger.Printf("Output received %s %s %s\n", uuid, stdout, sessionId)

	if uuid == "" {
		http.Error(w, "UUID is required", http.StatusBadRequest)
		return
	}

	if err := s.writeMessage(sessionId, []byte(stdout)); err != nil {
		s.logger.Errorf("Error writing message: %v", err)
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, r, http.StatusOK, nil)
}

var connections = make(map[string]*websocket.Conn)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (s *server) websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}
	defer conn.Close()
	sessionId := uuid.NewString()

	s.logger.Infof("New connection for session ID: %s", sessionId)

	connections[sessionId] = conn

	defer delete(connections, sessionId)

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		var submission Submission
		err = json.Unmarshal(p, &submission)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		submission.Uuid = uuid.NewString()
		submission.SessionId = sessionId

		if err := s.publishMessageQueue(submission); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
	}
}

func (s *server) writeMessage(sessionId string, message []byte) error {
	conn, ok := connections[sessionId]
	if !ok {
		return fmt.Errorf("no connection found for session ID %s", sessionId)
	}

	return conn.WriteMessage(websocket.TextMessage, message)
}
