package main

import (
	"github.com/nitinkumar-na/rabbitmq-n-go/consumer"
	"github.com/nitinkumar-na/rabbitmq-n-go/producer"
)

func main() {
	producer.PublishMessage()
	consumer.ListenToMessages()
}
