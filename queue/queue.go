package queue

import (
	"errors"

	"github.com/streadway/amqp"
)

type Queue struct {
	queue amqp.Queue

	Name             string `binding:"required"`
	Durable          bool
	DeleteWhenUnused bool
	Exclusive        bool
	NoWait           bool
	Args             amqp.Table
	Channel          *amqp.Channel

	bindOptions *BindOptions
	consumer    *Consumer
}

func New(Channel *amqp.Channel) *Queue {
	return &Queue{
		Durable:          false,
		DeleteWhenUnused: false,
		Exclusive:        false,
		NoWait:           false,
		Channel:          Channel,
		bindOptions:      &BindOptions{},
	}
}

func (config *Queue) From(queue *Queue) *Queue {
	config.Name = queue.Name
	config.Durable = queue.Durable
	config.DeleteWhenUnused = queue.DeleteWhenUnused
	config.Exclusive = queue.Exclusive
	config.NoWait = queue.NoWait
	config.Args = queue.Args
	config.Channel = queue.Channel
	config.bindOptions = queue.bindOptions
	return config
}

func (config *Queue) WithChannel(Channel *amqp.Channel) *Queue {
	config.Channel = Channel
	return config
}

func (config *Queue) WithName(Name string) *Queue {
	config.Name = Name
	return config
}

func (config *Queue) WithDurable(Durable bool) *Queue {
	config.Durable = Durable
	return config
}

func (config *Queue) WithDeleteWhenUnused(DeleteWhenUnused bool) *Queue {
	config.DeleteWhenUnused = DeleteWhenUnused
	return config
}

func (config *Queue) WithExclusive(Exclusive bool) *Queue {
	config.Exclusive = Exclusive
	return config
}

func (config *Queue) WithNoWait(NoWait bool) *Queue {
	config.NoWait = NoWait
	return config
}

func (config *Queue) WithArgs(Args amqp.Table) *Queue {
	config.Args = Args
	return config
}

func (config *Queue) Raw() amqp.Queue {
	return config.queue
}

func (config *Queue) Declare() (*Queue, error) {
	if config.Channel == nil {
		return nil, errors.New("channel is nil")
	}

	q, err := config.Channel.QueueDeclare(
		config.Name,
		config.Durable,
		config.DeleteWhenUnused,
		config.Exclusive,
		config.NoWait,
		config.Args,
	)

	config.bindOptions = NewBind().
		WithChannel(config.Channel).
		WithQueue(config)

	config.consumer = NewConsumer().
		WithQueue(config)

	if err != nil {
		return nil, err
	}

	config.queue = q
	return config, nil
}

func (config *Queue) Binding() *BindOptions {
	return config.bindOptions
}

func (config *Queue) Consumer() *Consumer {
	return config.consumer
}
