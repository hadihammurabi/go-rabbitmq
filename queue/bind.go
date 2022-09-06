package queue

import (
	"errors"

	"github.com/streadway/amqp"
)

type BindOptions struct {
	Name       string
	RoutingKey string
	Exchange   string
	NoWait     bool
	Args       amqp.Table
	Channel    *amqp.Channel
}

func NewBind() *BindOptions {
	return &BindOptions{}
}

func (config *BindOptions) WithChannel(Channel *amqp.Channel) *BindOptions {
	config.Channel = Channel
	return config
}

func (config *BindOptions) WithQueue(Queue *Queue) *BindOptions {
	config.Name = Queue.Name
	return config
}

func (config *BindOptions) WithRoutingKey(RoutingKey string) *BindOptions {
	config.RoutingKey = RoutingKey
	return config
}

func (config *BindOptions) WithExchange(Exchange string) *BindOptions {
	config.Exchange = Exchange
	return config
}

func (config *BindOptions) WithNoWait(NoWait bool) *BindOptions {
	config.NoWait = NoWait
	return config
}

func (config *BindOptions) WithArgs(Args amqp.Table) *BindOptions {
	config.Args = Args
	return config
}

func (config *BindOptions) Bind() error {
	if config.Channel == nil {
		return errors.New("channel is nil")
	}

	if config.Name == "" {
		return errors.New("queue name is empty")
	}

	return config.Channel.QueueBind(
		config.Name,
		config.RoutingKey,
		config.Exchange,
		config.NoWait,
		config.Args,
	)
}
