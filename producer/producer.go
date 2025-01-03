package producer

import (
	"context"
	"github.com/nitinkumar-na/rabbitmq-n-go/helper"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func PublishMessage(message string) {

	// start the publisher
	publisher := helper.StartRabbitMqClient("Producer")

	defer publisher.Conn.Close()
	defer publisher.Channel.Close()
	// TODO: this didn't work, read why?
	//defer helper.HandleError(publisher.conn.Close(), "Failed to close connection properly")
	//defer helper.HandleError(publisher.channel.Close(), "Failed to close channel properly")

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	err := publisher.Channel.PublishWithContext(ctx,
		"",
		publisher.Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	helper.HandleError(err, "Failed to publish the message")

	log.Printf(" [x] Sent %s\n", message)
}
