package apiserver

import (
	"context"
	"encoding/json"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (s *server) publishMessageQueue(submission Submission) error {
	conn := s.messageBroker.Get().(*amqp.Connection)
	defer s.messageBroker.Put(conn)

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	// takes code and userID and publishes it to the queue
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	myJson, err := json.Marshal(submission)
	if err != nil {
		return err
	}

	s.logger.Infof("Message broker received: %s", string(myJson))

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        myJson,
		})
	if err != nil {
		return err
	}
	s.logger.Infof(" [x] Sent %d %s %s\n", submission.UserId, submission.Code, submission.Stdin)
	return nil
}
