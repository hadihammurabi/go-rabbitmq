package queue

import (
	"errors"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	Queue     *Queue
	Consumer  string
	AutoACK   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      amqp.Table
}

func NewConsumer() *Consumer {
	return &Consumer{}
}

func (c *Consumer) WithQueue(Queue *Queue) *Consumer {
	c.Queue = Queue
	return c
}

func (c *Consumer) WithAutoACK(AutoACK bool) *Consumer {
	c.AutoACK = AutoACK
	return c
}

func (c *Consumer) WithExclusive(Exclusive bool) *Consumer {
	c.Exclusive = Exclusive
	return c
}

func (c *Consumer) WithNoLocal(NoLocal bool) *Consumer {
	c.NoLocal = NoLocal
	return c
}

func (c *Consumer) WithNoWait(NoWait bool) *Consumer {
	c.NoWait = NoWait
	return c
}

func (c *Consumer) WithArgs(Args amqp.Table) *Consumer {
	c.Args = Args
	return c
}

func (c Consumer) Consume() (<-chan amqp.Delivery, error) {
	if c.Queue == nil {
		return nil, errors.New("queue is nil")
	}

	if c.Queue.Channel == nil {
		return nil, errors.New("channel is nil")
	}

	consumer, err := c.Queue.Channel.Consume(
		c.Queue.Name,
		c.Consumer,
		c.AutoACK,
		c.Exclusive,
		c.NoLocal,
		c.NoWait,
		c.Args,
	)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}
