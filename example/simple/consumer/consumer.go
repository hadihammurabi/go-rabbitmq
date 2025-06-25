package main

// this is an example of consuming message from queue
// steps to do here are same as https://www.rabbitmq.com/tutorials/tutorial-one-go.html
// but using the wrapper API

import (
	"log"

	rabbitmq "github.com/hadihammurabi/go-rabbitmq"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	mq, err := rabbitmq.New("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to create a MQ")
	defer mq.Close()

	q, err := mq.Queue().
		WithName("hello").
		WithChannel(mq.Channel()).
		Declare()
	failOnError(err, "Failed to declare a queue")

	log.Println(" [*] Waiting for messages. To exit press CTRL+C")

	results, err := q.Consumer().Consume()
	failOnError(err, "Failed to register a consumer")

	for result := range results {
		log.Println("result: ", string(result.Body))
		err := result.Ack(false)
		if err != nil {
			continue
		}
	}
	forever := make(chan bool)
	<-forever
}
