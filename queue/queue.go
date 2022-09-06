package queue

import "github.com/streadway/amqp"

type Queue struct {
	queue amqp.Queue

	Name             string `binding:"required"`
	Durable          bool
	DeleteWhenUnused bool
	Exclusive        bool
	NoWait           bool
	Args             amqp.Table
	Channel          *amqp.Channel
}

func New() *Queue {
	return &Queue{}
}

func (config *Queue) From(queue *Queue) *Queue {
	config.Name = queue.Name
	config.Durable = queue.Durable
	config.DeleteWhenUnused = queue.DeleteWhenUnused
	config.Exclusive = queue.Exclusive
	config.NoWait = queue.NoWait
	config.Args = queue.Args
	config.Channel = queue.Channel
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

func (config *Queue) Declare() (*Queue, error) {
	q, err := config.Channel.QueueDeclare(
		config.Name,
		config.Durable,
		config.DeleteWhenUnused,
		config.Exclusive,
		config.NoWait,
		config.Args,
	)

	if err != nil {
		return nil, err
	}

	config.queue = q
	return config, nil
}
