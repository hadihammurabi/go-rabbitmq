package gorabbitmq

import (
	"github.com/hadihammurabi/go-rabbitmq/connection"
	"github.com/hadihammurabi/go-rabbitmq/exchange"
	"github.com/hadihammurabi/go-rabbitmq/queue"
	amqp "github.com/rabbitmq/amqp091-go"
)

type MQ struct {
	connection *connection.Connection
	channel    *amqp.Channel
	queue      *queue.Queue
	exchange   *exchange.Exchange
}

func New(opts ...func(*NewOptions)) (*MQ, error) {
	opt := NewOptions{}
	opt.Apply(opts...)

	var conn *connection.Connection
	if opt.url != "" {
		c, err := connection.New(opt.url)
		if err != nil {
			return nil, err
		}
		conn = c
	} else if opt.connection != nil {
		conn = connection.From(opt.connection)
	} else if opt.amqp != nil {
		conn = connection.FromAMQP(opt.amqp)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &MQ{
		connection: conn,
		channel:    ch,
		queue:      queue.New(ch),
		exchange:   exchange.New(ch),
	}, nil
}

func (mq *MQ) Connection() *connection.Connection {
	return mq.connection
}

func (mq *MQ) Channel() *amqp.Channel {
	return mq.channel
}

func (mq *MQ) Queue() *queue.Queue {
	return mq.queue
}

func (mq *MQ) Exchange() *exchange.Exchange {
	return mq.exchange
}

func (mq *MQ) Publish(publish *MQConfigPublish) error {
	return mq.channel.Publish(
		publish.Exchange,
		publish.RoutingKey,
		publish.Mandatory,
		publish.Immediate,
		publish.Message,
	)
}

func (mq *MQ) Close() {
	mq.channel.Close()
	mq.connection.Close()
}
