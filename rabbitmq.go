package gorabbitmq

import "github.com/streadway/amqp"

type MQ interface {
	GetConnection() *amqp.Connection
	GetChannel() *amqp.Channel
	GetQueue() amqp.Queue
	DeclareQueue(*MQConfigQueue) (amqp.Queue, error)
	DeclareExchange(*MQConfigExchange) error
	Publish(*MQConfigPublish) error
	Consume(amqp.Queue, *MQConfigConsume) (<-chan amqp.Delivery, error)
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
}

type MQConfigPublish struct {
	Exchange   string
	RoutingKey string
	Mandatory  bool
	Immediate  bool
	Message    amqp.Publishing
}
