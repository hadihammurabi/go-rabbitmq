package gorabbitmq

import (
	"github.com/streadway/amqp"

	queue "github.com/hadihammurabi/go-rabbitmq/queue"
)

type MQ struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      *queue.Queue
}

func New(url string) (*MQ, error) {
	conn, err := NewConnection(NewConnectionOptions().SetURL(url))
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &MQ{
		connection: conn,
		channel:    ch,
	}, nil
}

func NewFromConnection(conn *amqp.Connection) (*MQ, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &MQ{
		connection: conn,
		channel:    ch,
	}, nil
}

func (mq *MQ) Connection() *amqp.Connection {
	return mq.connection
}

func (mq *MQ) Channel() *amqp.Channel {
	return mq.channel
}

func (mq *MQ) Queue() *queue.Queue {
	return mq.queue
}

func (mq *MQ) QueueDeclare(config *queue.Queue) (*queue.Queue, error) {
	q, err := queue.New().From(config).Declare()
	if err != nil {
		return nil, err
	}

	mq.queue = q
	return mq.queue, nil
}

func (mq *MQ) QueueBind(config *MQConfigQueueBind) error {
	err := NewQueueBind(mq.channel, config)
	if err != nil {
		return err
	}

	return nil
}

func (mq *MQ) ExchangeDeclare(config *MQConfigExchange) error {
	err := NewExchange(mq.channel, config)
	if err != nil {
		return err
	}

	return nil
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

func (mq *MQ) Consume(consume *MQConfigConsume) (<-chan amqp.Delivery, error) {
	if consume == nil {
		consume = &MQConfigConsume{}
	}

	consumer, err := mq.channel.Consume(
		mq.queue.Name,
		consume.Consumer,
		consume.AutoACK,
		consume.Exclusive,
		consume.NoLocal,
		consume.NoWait,
		consume.Args,
	)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}

func (mq *MQ) Close() {
	mq.channel.Close()
	mq.connection.Close()
}
