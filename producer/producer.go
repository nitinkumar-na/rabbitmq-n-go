package producer

import (
	"context"
	"fmt"
	"github.com/nitinkumar-na/rabbitmq-with-go/helper"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func PublishMessage() {
	fmt.Println("Producer Started")

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
	helper.HandleError(err, "Failed to create q queue")

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	message := "Hello From Producer"

	err = ch.PublishWithContext(ctx,
		"",
		queue.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	helper.HandleError(err, "Failed to publish the message")

	log.Printf(" [x] Sent %s\n", message)
}
