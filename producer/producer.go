package producer

import (
	"context"
	"fmt"
	"github.com/nitinkumar-na/rabbitmq-n-go/helper"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type Publisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func PublishMessage(message string) {

	// validate user input
	if helper.IsBlank(message) {
		log.Printf("Message can't be blank")
		return
	}

	// start the publisher
	publisher := startPublisher()

	defer publisher.conn.Close()
	defer publisher.channel.Close()
	// TODO: this didn't work, read why?
	//defer helper.HandleError(publisher.conn.Close(), "Failed to close connection properly")
	//defer helper.HandleError(publisher.channel.Close(), "Failed to close channel properly")

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	err := publisher.channel.PublishWithContext(ctx,
		"",
		publisher.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	helper.HandleError(err, "Failed to publish the message")

	log.Printf(" [x] Sent %s\n", message)
}

func startPublisher() Publisher {
	fmt.Println("Produce Starting...")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	helper.HandleError(err, "Failed to connect to RabbitMQ")
	//defer conn.Close()

	ch, err := conn.Channel()
	helper.HandleError(err, "Failed to open a channel")
	//defer ch.Close()

	queue, err := ch.QueueDeclare(
		"my-queue",
		false,
		false,
		false,
		false,
		nil,
	)
	helper.HandleError(err, "Failed to create q queue")

	fmt.Println("Producer Started Successfully")

	return Publisher{conn, ch, queue}
}
