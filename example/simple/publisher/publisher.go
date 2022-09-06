package main

// this is an example of publising message directly into queue
// steps to do here are same as https://www.rabbitmq.com/tutorials/tutorial-one-go.html
// but using the wrapper API

import (
	"fmt"
	"log"
	"sync"

	rabbitmq "github.com/hadihammurabi/go-rabbitmq"
	"github.com/hadihammurabi/go-rabbitmq/exchange"
	"github.com/streadway/amqp"
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

	err = mq.Exchange().
		WithName("hello").
		WithType(exchange.TypeDirect).
		Declare()
	failOnError(err, "Failed to create a channel")

	q, err := mq.Queue().
		WithName("hello").
		Declare()
	failOnError(err, "Failed to create a queue")

	err = q.Binding().
		WithExchange("hello").
		Bind()
	failOnError(err, "Failed to bind queue")

	var wg sync.WaitGroup
	max := 10
	wg.Add(max)
	for i := 0; i < max; i++ {
		go func(a int) {
			defer wg.Done()
			body := fmt.Sprintf("Hello World %d !", a)
			err = mq.Publish(&rabbitmq.MQConfigPublish{
				RoutingKey: mq.Queue().Name,
				Message: amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(body),
				},
			})
		}(i)
	}
	wg.Wait()
}
