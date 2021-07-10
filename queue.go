package gorabbitmq

import "github.com/streadway/amqp"

type MQConfigQueue struct {
	Name             string `binding:"required"`
	Durable          bool
	DeleteWhenUnused bool
	Exclusive        bool
	NoWait           bool
	Args             amqp.Table
}

func NewQueueOptions() *MQConfigQueue {
	return &MQConfigQueue{}
}

func (config *MQConfigQueue) SetName(Name string) *MQConfigQueue {
	config.Name = Name
	return config
}

func (config *MQConfigQueue) SetDurable(Durable bool) *MQConfigQueue {
	config.Durable = Durable
	return config
}

func (config *MQConfigQueue) SetDeleteWhenUnused(DeleteWhenUnused bool) *MQConfigQueue {
	config.DeleteWhenUnused = DeleteWhenUnused
	return config
}

func (config *MQConfigQueue) SetExclusive(Exclusive bool) *MQConfigQueue {
	config.Exclusive = Exclusive
	return config
}

func (config *MQConfigQueue) SetNoWait(NoWait bool) *MQConfigQueue {
	config.NoWait = NoWait
	return config
}

func (config *MQConfigQueue) SetArgs(Args amqp.Table) *MQConfigQueue {
	config.Args = Args
	return config
}

func NewQueue(channel *amqp.Channel, queue *MQConfigQueue) (amqp.Queue, error) {
	return channel.QueueDeclare(
		queue.Name,
		queue.Durable,
		queue.DeleteWhenUnused,
		queue.Exclusive,
		queue.NoWait,
		queue.Args,
	)
}
