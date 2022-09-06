package main

// this is an example of consuming message from queue
// steps to do here are same as https://www.rabbitmq.com/tutorials/tutorial-one-go.html
// but using the wrapper API

import (
	"fmt"
	"log"

	rabbitmq "github.com/hadihammurabi/go-rabbitmq"
	queue "github.com/hadihammurabi/go-rabbitmq/queue"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	def, err := rabbitmq.New("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to create a MQ")
	defer def.Close()

	del, err := rabbitmq.NewFromConnection(def.Connection())
	failOnError(err, fmt.Sprintf("%v", err))
	defer del.Close()

	_, err = del.QueueDeclare(&queue.MQConfigQueue{
		Name: "hello",
	})
	failOnError(err, "Failed to declare a queue")

	log.Println(" [*] Waiting for messages. To exit press CTRL+C")

	results, err := del.Consume(nil)
	failOnError(err, "Failed to register a consumer")

	for result := range results {
		log.Println("result: ", string(result.Body))
		result.Ack(false)
	}
	forever := make(chan bool)
	<-forever
}
