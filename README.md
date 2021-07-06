# go-rabbitmq
Golang AMQP wrapper for RabbitMQ with better API 

# Table of Contents
* [Background](#background)
* [Features](#features)
* [Usage](#usage)
  * [Installation](#installation)
  * [Connect to RabbitMQ](#connect-to-rabbitmq)
  * [Declare Queue](#declare-queue)
  * [Declare Exchange](#declare-exchange)
  * [Bind Queue to Exchange](#bind-queue-to-exchange)
  * [Publish a Message](#publish-a-message)
  * [Consume Messages](#consume-messages)
* [How It Works](#how-it-works)
* [License](#license)
* [Contributing](#contributing)

# Background
In [Golang](https://golang.org), to use [RabbitMQ](https://www.rabbitmq.com) with [AMQP](https://www.amqp.org) has advantages, especially in messaging systems.
It's done with the [AMQP connector](https://github.com/streadway/amqp).
But, the problem is it has a less convenient API.
Programmers have to write something that should be set by default.
For example, when creating a queue on RabbitMQ.
We have to do this.
```go
q, err := ch.QueueDeclare(
  "hello", // name
  false,   // durable
  false,   // delete when unused
  false,   // exclusive
  false,   // no-wait
  nil,     // arguments
)
failOnError(err, "Failed to declare a queue")
```
Too many `false` in there, which should be set as the default value.

By using [this](.) module, we can do same think with less code. See above.
```go
q, err := mq.NewQueue(
  ch,
  &gorabbitmq.MQConfigQueue{
   Name: "hello",
  },
)
failOnError(err, "Failed to declare a queue")
```
No need to write `false`, because it is the default value.

So, to conclude, this module makes it easy to use amqp for rabbitmq.

# Features
* Built on top of famous [AMQP connector](https://github.com/streadway/amqp) in Go.
* All object reference like connection, channel, queue, etc are original by the AMQP connector (no monkey patch or any modifications).
* It has construction API and API with builder pattern.
* Does not modify incoming messages, so it can be controlled manually.

# Usage
## Installation
Inside terminal emulator, simply run command below.
```sh
go get github.com/hadihammurabi/go-rabbitmq
```

After installation it can be imported into any Go project.
For example.
```go
package main

import (
 rabbitmq "github.com/hadihammurabi/go-rabbitmq"
)
```

## Connect to RabbitMQ
It can do as below.
```go
mq, err := rabbitmq.NewMQ(&rabbitmq.MQConfigConnection{
	URL: "amqp://guest:guest@localhost:5672/",
})
if err != nil {
 log.Fatal(err)
}

// don't forget to close the connection and channel
defer func() {
 mq.GetConnection().Close()
 mq.GetChannel().Close()
}()
```

## Declare Queue
Queue declaration can be done like this, after connecting to mq of course.
> It only connects to the queue if the queue exists or create one if it doesn't exist. (RabbitMQ behavior)
```go
q, err := rabbitmq.NewQueue(mq.GetChannel(), &rabbitmq.MQConfigQueue{
  Name: "hello",
})
if err != nil {
 log.Fatal(err)
}
```

## Declare Exchange
Exchange declaration can be done like this, after connecting to mq of course.
```go
err := rabbitmq.NewExchange(mq.GetChannel(), &rabbitmq.MQConfigExchange{
  Name: "hello",
  Type: rabbitmq.ExchangeTypeFanout,
})
if err != nil {
 log.Fatal(err)
}
```

## Bind Queue to Exchange
Every message published to exchange will be distributed to every bound queue.
To bind queue with exchange, follow example below.
```go
err := rabbitmq.NewQueueBind(mq.GetChannel(), &rabbitmq.MQConfigBind{
  Name:     q.Name,
  Exchange: "hello",
})
if err != nil {
 log.Fatal(err)
}
```

## Publish a Message
A message can be sent to exchange by mentioning the name exchange.
Publishing a message can do like this.
```go
err := mq.Publish(&rabbitmq.MQConfigPublish{
  Exchange: "hello",
  Message: amqp.Publishing{
    ContentType: "text/plain",
    Body:        []byte(body),
  },
})
if err != nil {
 log.Fatal(err)
}
```

## Consume Messages
Every message in the queue can be consumed by the queue consumer.
To consume messages in queue can do like this.

> The following code will run forever to listen for new message in queue.
```go
msgs, err := mq.Consume(q, &rabbitmq.MQConfigConsume{})
if err != nil {
 log.Fatal(err)
}

forever := make(chan bool)
go func() {
  for msg := range msgs {
    fmt.Println(string(msg.Body))
    msg.Ack(false)
  }
}()
<-forever
```

# How It Works
# License
This project is under [Mozilla Public License 2.0](./LICENSE).

# Contributing
Realy love any contribution. Feel free to create a [Pull Request](https://docs.github.com/en/github/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/creating-a-pull-request) with following this [commit convention](https://github.com/angular/angular/blob/master/CONTRIBUTING.md#commit). 
