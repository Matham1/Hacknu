package apiserver

import (
	"log"
	"net/http"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Start(config AppConfig) error {
	log.Printf("Connecting to RabbitMQ: %s\n", config.RabbitMQ)
	var rabbitMQPool = &sync.Pool{
		New: func() interface{} {
			conn, err := amqp.Dial(config.RabbitMQ)
			if err != nil {
				log.Fatalf("Failed to connect to RabbitMQ: %v", err)
				return nil
			}
			log.Print("Connected to RabbitMQ")
			return conn
		},
	}
	srv := newServer(rabbitMQPool)
	log.Printf("Server is listening on port %s...\n", config.BindAddr)
	return http.ListenAndServe(config.BindAddr, srv)
}
