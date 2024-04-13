package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
}

type Submission struct {
	UserId    int    `json:"userId"`
	Code      string `json:"code"`
	Stdin     string `json:"stdin"`
	Uuid      string `json:"uuid"`
	SessionId string `json:"sessionId"`
}

// Constructor of new server
func newServer() *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
	}
	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/read", homeHandler).Methods("GET")
}

func (s *server) respond(w http.ResponseWriter, _ *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.logger.Errorf("Code: %d, Error: %v", code, err)
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}
