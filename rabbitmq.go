package gorabbitmq

import "github.com/streadway/amqp"

type MQ interface {
	GetConnection() *amqp.Connection
	GetChannel(name ...string) *amqp.Channel
	GetQueue() amqp.Queue
	DeclareQueue(string, *MQConfigQueue) (amqp.Queue, error)
	DeclareExchange(string, *MQConfigExchange) error
	QueueBind(string, *MQConfigBind) error
	Publish(string, *MQConfigPublish) error
	Consume(string, *MQConfigConsume) (<-chan amqp.Delivery, error)
	Close()
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
