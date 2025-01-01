package consumer

import (
	"fmt"
	"github.com/nitinkumar-na/rabbitmq-n-go/helper"
	"github.com/rabbitmq/amqp091-go"
	"log"
)

func ListenToMessages() {
	fmt.Println("Consumer Started")

	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	helper.HandleError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	helper.HandleError(err, "Failed to open a channel")
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"my-queue",
		false,
		false,
		false,
		false,
		nil,
	)
	helper.HandleError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	helper.HandleError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
}

func StartConsumer() {

}
