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
# Usage
## Installation
## Connect to RabbitMQ
## Declare Queue
## Declare Exchange
## Bind Queue to Exchange
## Consume Messages
# How It Works
# License
This project is under [Mozilla Public License 2.0](./LICENSE).

# Contributing
Realy love any contribution. Feel free to create a [Pull Request](https://docs.github.com/en/github/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/creating-a-pull-request) with following this [commit convention](https://github.com/angular/angular/blob/master/CONTRIBUTING.md#commit). 
