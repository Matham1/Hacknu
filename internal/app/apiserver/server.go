package apiserver

import (
	"encoding/json"
	"net/http"

	db "github.com/abd-rakhman/qysqa-back/internal/db/sqlc"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router   *mux.Router
	logger   *logrus.Logger
	database *db.Queries
}

// Constructor of new server
func newServer(database *db.Queries) *server {
	s := &server{
		router:   mux.NewRouter(),
		logger:   logrus.New(),
		database: database,
	}
	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/", homeHandler).Methods("GET")
	s.router.HandleFunc("/create/", s.createContest).Methods("POST")
	s.router.HandleFunc("/last-contests/", s.getLastContestsHandler).Methods("GET")
	s.router.HandleFunc("/contest/{id}/", s.getContestById).Methods("GET")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
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
