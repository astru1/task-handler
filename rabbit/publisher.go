package local_rabbit

import (
	"awesomeProject/database"
	"encoding/json"
	"fmt"
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

func main() {
	fmt.Println("Starting up")
	conn, err := amqp.Dial("amqp://guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to RabbitMQ Instance")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(q)

	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/Plain",
			Body:        []byte("Hello World"),
		},
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Successfully published message to queue")

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
