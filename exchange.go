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

func NewExchangeOptions() *MQConfigExchange {
	return &MQConfigExchange{}
}

func (config *MQConfigExchange) SetName(Name string) *MQConfigExchange {
	config.Name = Name
	return config
}

func (config *MQConfigExchange) SetType(Type ExchangeType) *MQConfigExchange {
	config.Type = Type
	return config
}

func (config *MQConfigExchange) SetDurable(Durable bool) *MQConfigExchange {
	config.Durable = Durable
	return config
}

func (config *MQConfigExchange) SetAutoDeleted(AutoDeleted bool) *MQConfigExchange {
	config.AutoDeleted = AutoDeleted
	return config
}

func (config *MQConfigExchange) SetInternal(Internal bool) *MQConfigExchange {
	config.Internal = Internal
	return config
}

func (config *MQConfigExchange) SetNoWait(NoWait bool) *MQConfigExchange {
	config.NoWait = NoWait
	return config
}

func (config *MQConfigExchange) SetArgs(Args amqp.Table) *MQConfigExchange {
	config.Args = Args
	return config
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
