package gorabbitmq

import "github.com/streadway/amqp"

const (
	ChannelDefault string = "default"
)

type MQ interface {
	Connection() *amqp.Connection
	Channel() *amqp.Channel
	QueueDeclare(*MQConfigQueue) (amqp.Queue, error)
	Queue() amqp.Queue
	QueueBind(*MQConfigQueueBind) error
	ExchangeDeclare(*MQConfigExchange) error
	Publish(*MQConfigPublish) error
	Consume(*MQConfigConsume) (<-chan amqp.Delivery, error)
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
