package apiserver

import (
	"log"
	"net/http"
)

func Start(config AppConfig) error {
	log.Printf("Connecting to RabbitMQ: %s\n", config.RabbitMQ)
	srv := newServer()
	log.Printf("Server is listening on port %s...\n", config.BindAddr)
	return http.ListenAndServe(config.BindAddr, srv)
}
