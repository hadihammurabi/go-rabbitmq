package gorabbitmq

import "github.com/streadway/amqp"

type ExchangeType string

const (
	ExchangeTypeDirect ExchangeType = "direct"
	ExchangeTypeTopic  ExchangeType = "topic"
	ExchangeTypeFanout ExchangeType = "fanout"
)

type MQConfigExchange struct {
	Name        string
	Type        ExchangeType
	Durable     bool
	AutoDeleted bool
	Internal    bool
	NoWait      bool
	Args        amqp.Table
}

func NewExchange(channel *amqp.Channel, exchange *MQConfigExchange) error {
	return channel.ExchangeDeclare(
		exchange.Name,
		string(exchange.Type),
		exchange.Durable,
		exchange.AutoDeleted,
		exchange.Internal,
		exchange.NoWait,
		exchange.Args,
	)
}
