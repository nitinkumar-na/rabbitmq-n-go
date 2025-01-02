package main

import (
	"fmt"
	"github.com/nitinkumar-na/rabbitmq-n-go/consumer"
	"github.com/nitinkumar-na/rabbitmq-n-go/helper"
	"github.com/nitinkumar-na/rabbitmq-n-go/producer"
)

// 1. Fire and Forget
//func main() {
//	producer.PublishMessage()
//	consumer.ListenToMessages()
//}

// 2. Take user input, publish and consumer prints
func main() {

	for {
		input := getUserInput()

		if input == "exit" {
			break
		}

		producer.PublishMessage(input)
		consumer.ListenToMessages()
	}
}

func getUserInput() string {
	fmt.Println("Enter your message, or type 'exit' to quit")
	var input string
	fmt.Scan(&input)

	if helper.IsBlank(input) {
		fmt.Println("Message can not be empty")
		getUserInput()
	}

	return input
}
