package queue

import "github.com/streadway/amqp"

type BindOptions struct {
	Name       string
	RoutingKey string
	Exchange   string
	NoWait     bool
	Args       amqp.Table
	Channel    *amqp.Channel
}

func NewBind(channel *amqp.Channel, queue *Queue) *BindOptions {
	return &BindOptions{
		Channel: channel,
		Name:    queue.Name,
	}
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
	return config.Channel.QueueBind(
		config.Name,
		config.RoutingKey,
		config.Exchange,
		config.NoWait,
		config.Args,
	)
}
