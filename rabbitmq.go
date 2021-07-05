package gorabbitmq

import "github.com/streadway/amqp"

type MQ interface {
	GetConnection() *amqp.Connection
	GetChannel() *amqp.Channel
	GetQueue() amqp.Queue
	Publish(publish *MQConfigPublish) error
	Consume(queue amqp.Queue, consume *MQConfigConsume) error
}

type MQConfig struct {
	Connection *MQConfigConnection
	Exchange   *MQConfigExchange
}

type MQConfigConsume struct {
	Name      string
	Consumer  string
	AutoACK   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      amqp.Table
	OnMessage func(msgs <-chan amqp.Delivery)
}

type MQConfigPublish struct {
	Exchange   string
	RoutingKey string
	Mandatory  bool
	Immediate  bool
	Message    amqp.Publishing
}
