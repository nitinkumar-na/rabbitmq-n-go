package consumer

import (
	"github.com/nitinkumar-na/rabbitmq-n-go/helper"
	"log"
)

func ListenToMessages() {
	client := helper.StartRabbitMqClient("Consumer")

	msgs := helper.ConsumerMessages(&client)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
}
