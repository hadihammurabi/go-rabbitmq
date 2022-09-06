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
	mq, err := rabbitmq.New("amqp://guest:guest@localhost:5672/")
	failOnError(err, fmt.Sprintf("%v", err))
	defer mq.Close()

	err = mq.Exchange().
		WithName("hello").
		WithType(exchange.TypeDirect).
		WithChannel(mq.Channel()).
		Declare()
	failOnError(err, fmt.Sprintf("%v", err))

	q, err := mq.Queue().
		WithName("hello").
		WithChannel(mq.Channel()).
		Declare()
	failOnError(err, fmt.Sprintf("%v", err))

	err = q.Binding().
		WithExchange("hello").
		Bind()
	failOnError(err, fmt.Sprintf("%v", err))

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
