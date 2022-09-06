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
	mq, err := rabbitmq.New("amqp://formiadmin:7shZA7HfDFAZ8WBa@rabbitmq-amqp.engine.159.223.41.6.sslip.io:30904/")
	failOnError(err, "Failed to create a MQ")
	defer mq.Close()

	_, err = mq.Queue().
		WithName("hello").
		WithChannel(mq.Channel()).
		Declare()
	failOnError(err, "Failed to declare a queue")

	log.Println(" [*] Waiting for messages. To exit press CTRL+C")

	results, err := mq.Consume(nil)
	failOnError(err, "Failed to register a consumer")

	for result := range results {
		log.Println("result: ", string(result.Body))
		result.Ack(false)
	}
	forever := make(chan bool)
	<-forever
}
