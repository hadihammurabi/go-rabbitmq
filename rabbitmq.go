package gorabbitmq

import "github.com/streadway/amqp"

const (
	ChannelDefault = "default"
)

type MQ interface {
	GetConnection() *amqp.Connection
	CreateChannel(name ...string) (*amqp.Channel, error)
	GetChannel(name ...string) *amqp.Channel
	WithChannel(name ...string) MQ
	DeclareQueue(*MQConfigQueue) (amqp.Queue, error)
	GetQueue() amqp.Queue
	QueueBind(*MQConfigBind) error
	DeclareExchange(*MQConfigExchange) error
	Publish(ch string, publish *MQConfigPublish) error
	Consume(ch string, queue amqp.Queue, consume *MQConfigConsume) (<-chan amqp.Delivery, error)
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
