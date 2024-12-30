package main

import (
	"github.com/nitinkumar-na/rabbitmq-with-go/consumer"
	"github.com/nitinkumar-na/rabbitmq-with-go/producer"
)

func main() {
	consumer.TestConsumer()
	producer.TestProducer()
}
