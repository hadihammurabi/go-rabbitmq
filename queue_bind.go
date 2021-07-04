package gorabbitmq

import "github.com/streadway/amqp"

type MQConfigBind struct {
	Name       string
	RoutingKey string
	Exchange   string
	NoWait     bool
	Args       amqp.Table
}

func NewQueueBind(channel *amqp.Channel, bind *MQConfigBind) error {
	return channel.QueueBind(
		bind.Name,
		bind.RoutingKey,
		bind.Exchange,
		bind.NoWait,
		bind.Args,
	)
}
