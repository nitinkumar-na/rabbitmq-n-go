package main

import (
	"github.com/nitinkumar-na/rabbitmq-with-go/consumer"
	"github.com/nitinkumar-na/rabbitmq-with-go/producer"
)

func main() {
	producer.PublishMessage()
	consumer.ListenToMessages()
}
