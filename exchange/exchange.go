package exchange

import (
	"errors"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Type string

const (
	TypeDirect Type = "direct"
	TypeTopic  Type = "topic"
	TypeFanout Type = "fanout"
)

type Exchange struct {
	Name        string
	Type        Type
	Durable     bool
	AutoDeleted bool
	Internal    bool
	NoWait      bool
	Args        amqp.Table
	Channel     *amqp.Channel
}

func New(Channel *amqp.Channel) *Exchange {
	return &Exchange{
		Channel: Channel,
	}
}

func (config *Exchange) From(exchange *Exchange) *Exchange {
	config.Name = exchange.Name
	config.Type = exchange.Type
	config.Durable = exchange.Durable
	config.AutoDeleted = exchange.AutoDeleted
	config.Internal = exchange.Internal
	config.NoWait = exchange.NoWait
	config.Args = exchange.Args
	config.Channel = exchange.Channel
	return config
}

func (config *Exchange) WithChannel(Channel *amqp.Channel) *Exchange {
	config.Channel = Channel
	return config
}

func (config *Exchange) WithName(Name string) *Exchange {
	config.Name = Name
	return config
}

func (config *Exchange) WithType(Type Type) *Exchange {
	config.Type = Type
	return config
}

func (config *Exchange) WithDurable(Durable bool) *Exchange {
	config.Durable = Durable
	return config
}

func (config *Exchange) WithAutoDeleted(AutoDeleted bool) *Exchange {
	config.AutoDeleted = AutoDeleted
	return config
}

func (config *Exchange) WithInternal(Internal bool) *Exchange {
	config.Internal = Internal
	return config
}

func (config *Exchange) WithNoWait(NoWait bool) *Exchange {
	config.NoWait = NoWait
	return config
}

func (config *Exchange) WithArgs(Args amqp.Table) *Exchange {
	config.Args = Args
	return config
}

func (config *Exchange) Declare() error {
	if config.Channel == nil {
		return errors.New("channel is nil")
	}

	return config.Channel.ExchangeDeclare(
		config.Name,
		string(config.Type),
		config.Durable,
		config.AutoDeleted,
		config.Internal,
		config.NoWait,
		config.Args,
	)
}
