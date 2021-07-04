package gorabbitmq

import "github.com/streadway/amqp"

type ConfigQueue struct {
	Name             string
	Durable          bool
	DeleteWhenUnused bool
	Exclusive        bool
	NoWait           bool
	Args             amqp.Table
}

func NewQueue(channel *amqp.Channel, queue *ConfigQueue) (amqp.Queue, error) {
	return channel.QueueDeclare(
		queue.Name,
		queue.Durable,
		queue.DeleteWhenUnused,
		queue.Exclusive,
		queue.NoWait,
		queue.Args,
	)
}
