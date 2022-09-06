package gorabbitmq

import "github.com/streadway/amqp"

type MQConfigQueue struct {
	Name             string `binding:"required"`
	Durable          bool
	DeleteWhenUnused bool
	Exclusive        bool
	NoWait           bool
	Args             amqp.Table
	Channel          *amqp.Channel
}

func Builder() *MQConfigQueue {
	return &MQConfigQueue{}
}

func (config *MQConfigQueue) From(queue *MQConfigQueue) *MQConfigQueue {
	config.Name = queue.Name
	config.Durable = queue.Durable
	config.DeleteWhenUnused = queue.DeleteWhenUnused
	config.Exclusive = queue.Exclusive
	config.NoWait = queue.NoWait
	config.Args = queue.Args
	config.Channel = queue.Channel
	return config
}

func (config *MQConfigQueue) WithChannel(Channel *amqp.Channel) *MQConfigQueue {
	config.Channel = Channel
	return config
}

func (config *MQConfigQueue) WithName(Name string) *MQConfigQueue {
	config.Name = Name
	return config
}

func (config *MQConfigQueue) WithDurable(Durable bool) *MQConfigQueue {
	config.Durable = Durable
	return config
}

func (config *MQConfigQueue) WithDeleteWhenUnused(DeleteWhenUnused bool) *MQConfigQueue {
	config.DeleteWhenUnused = DeleteWhenUnused
	return config
}

func (config *MQConfigQueue) WithExclusive(Exclusive bool) *MQConfigQueue {
	config.Exclusive = Exclusive
	return config
}

func (config *MQConfigQueue) WithNoWait(NoWait bool) *MQConfigQueue {
	config.NoWait = NoWait
	return config
}

func (config *MQConfigQueue) WithArgs(Args amqp.Table) *MQConfigQueue {
	config.Args = Args
	return config
}

func (config *MQConfigQueue) Build() (amqp.Queue, error) {
	return config.Channel.QueueDeclare(
		config.Name,
		config.Durable,
		config.DeleteWhenUnused,
		config.Exclusive,
		config.NoWait,
		config.Args,
	)
}
