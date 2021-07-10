package gorabbitmq

import "github.com/streadway/amqp"

type MQConfigQueueBind struct {
	Name       string
	RoutingKey string
	Exchange   string
	NoWait     bool
	Args       amqp.Table
}

func NewQueueBindOptions() *MQConfigQueueBind {
	return &MQConfigQueueBind{}
}

func (config *MQConfigQueueBind) SetName(Name string) *MQConfigQueueBind {
	config.Name = Name
	return config
}

func (config *MQConfigQueueBind) SetRoutingKey(RoutingKey string) *MQConfigQueueBind {
	config.RoutingKey = RoutingKey
	return config
}

func (config *MQConfigQueueBind) SetExchange(Exchange string) *MQConfigQueueBind {
	config.Exchange = Exchange
	return config
}

func (config *MQConfigQueueBind) SetNoWait(NoWait bool) *MQConfigQueueBind {
	config.NoWait = NoWait
	return config
}

func (config *MQConfigQueueBind) SetArgs(Args amqp.Table) *MQConfigQueueBind {
	config.Args = Args
	return config
}

func NewQueueBind(channel *amqp.Channel, bind *MQConfigQueueBind) error {
	return channel.QueueBind(
		bind.Name,
		bind.RoutingKey,
		bind.Exchange,
		bind.NoWait,
		bind.Args,
	)
}
