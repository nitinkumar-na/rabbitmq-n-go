package main

import (
	"github.com/nitinkumar-na/rabbitmq-n-go/consumer"
	"github.com/nitinkumar-na/rabbitmq-n-go/producer"
)

// 1. Fire and Forget
//func main() {
//	producer.PublishMessage()
//	consumer.ListenToMessages()
//}

// 2. Take user input, publish and consumer prints
func main() {
	producer.PublishMessage("Hello World")
	consumer.ListenToMessages()
}
