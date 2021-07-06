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
