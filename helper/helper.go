package helper

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"strings"
)

type RabbitMqConn struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   amqp.Queue
}

func HandleError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func IsBlank(str string) bool {
	str = strings.TrimSpace(str)
	return len(str) == 0
}

func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

func StartRabbitMqClient(client string) RabbitMqConn {
	fmt.Println(client, "Starting...")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	HandleError(err, client+" Failed to connect to RabbitMQ")
	//defer conn.Close()

	ch, err := conn.Channel()
	HandleError(err, client+" Failed to open a channel")
	//defer ch.Close()

	queue, err := ch.QueueDeclare(
		"my-queue",
		false,
		false,
		false,
		false,
		nil,
	)
	HandleError(err, client+" Failed to create q queue")

	fmt.Println(client, " Started Successfully")

	return RabbitMqConn{conn, ch, queue}
}

// ConsumerMessages returns a read only channel
func ConsumerMessages(client *RabbitMqConn) <-chan amqp.Delivery {
	msgs, err := client.Channel.Consume(
		client.Queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	HandleError(err, "Failed to register a consumer")

	return msgs
}

// Test GoLand
