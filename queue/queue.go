package gorabbitmq

import "github.com/streadway/amqp"

type Options struct {
	Name             string `binding:"required"`
	Durable          bool
	DeleteWhenUnused bool
	Exclusive        bool
	NoWait           bool
	Args             amqp.Table
	Channel          *amqp.Channel
}

func Builder() *Options {
	return &Options{}
}

func (config *Options) From(queue *Options) *Options {
	config.Name = queue.Name
	config.Durable = queue.Durable
	config.DeleteWhenUnused = queue.DeleteWhenUnused
	config.Exclusive = queue.Exclusive
	config.NoWait = queue.NoWait
	config.Args = queue.Args
	config.Channel = queue.Channel
	return config
}

func (config *Options) WithChannel(Channel *amqp.Channel) *Options {
	config.Channel = Channel
	return config
}

func (config *Options) WithName(Name string) *Options {
	config.Name = Name
	return config
}

func (config *Options) WithDurable(Durable bool) *Options {
	config.Durable = Durable
	return config
}

func (config *Options) WithDeleteWhenUnused(DeleteWhenUnused bool) *Options {
	config.DeleteWhenUnused = DeleteWhenUnused
	return config
}

func (config *Options) WithExclusive(Exclusive bool) *Options {
	config.Exclusive = Exclusive
	return config
}

func (config *Options) WithNoWait(NoWait bool) *Options {
	config.NoWait = NoWait
	return config
}

func (config *Options) WithArgs(Args amqp.Table) *Options {
	config.Args = Args
	return config
}

func (config *Options) Build() (amqp.Queue, error) {
	return config.Channel.QueueDeclare(
		config.Name,
		config.Durable,
		config.DeleteWhenUnused,
		config.Exclusive,
		config.NoWait,
		config.Args,
	)
}
