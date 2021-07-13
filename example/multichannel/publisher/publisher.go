package main

// this is an example of publising message directly into queue with multi channel
// steps to do here are same as https://www.rabbitmq.com/tutorials/tutorial-one-go.html
// but using the wrapper API

import (
	"fmt"
	"log"
	"sync"

	rabbitmq "github.com/hadihammurabi/go-rabbitmq"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	def, err := rabbitmq.NewMQ("amqp://guest:guest@localhost:5672/")
	failOnError(err, fmt.Sprintf("%v", err))
	defer def.Close()

	del, err := rabbitmq.NewMQFromConnection(def.Connection())
	failOnError(err, fmt.Sprintf("%v", err))
	defer del.Close()

	_, err = del.QueueDeclare(rabbitmq.NewQueueOptions().SetName("hello"))
	failOnError(err, fmt.Sprintf("%v", err))

	var wg sync.WaitGroup
	max := 10
	wg.Add(max)
	for i := 0; i < max; i++ {
		go func(a int) {
			defer wg.Done()
			body := fmt.Sprintf("Hello World %d !", a)
			err = del.Publish(&rabbitmq.MQConfigPublish{
				RoutingKey: del.Queue().Name,
				Message: amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(body),
				},
			})
		}(i)
	}
	wg.Wait()
}
