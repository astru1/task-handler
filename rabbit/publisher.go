package local_rabbit

import (
	"awesomeProject/database"
	"encoding/json"
	"github.com/streadway/amqp"
)

type LocalQueue struct {
	queueName string
	ch        *amqp.Channel
}

func InitQueue(connectionString, queueName string) (*LocalQueue, error) {
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}
	return &LocalQueue{
		queueName: queueName,
		ch:        ch,
	}, nil
}
func (lq *LocalQueue) AddToQueue(task database.Task) error {
	marshaledTask, err := json.Marshal(task)
	if err != nil {
		return err
	}
	return lq.ch.Publish(
		"",
		lq.queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        marshaledTask,
		},
	)
}
