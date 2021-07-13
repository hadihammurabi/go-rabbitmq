package gorabbitmq

import (
	"github.com/streadway/amqp"
)

type mqDefault struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      amqp.Queue
}

func NewMQ(url string) (MQ, error) {
	conn, err := NewConnection(NewConnectionOptions().SetURL(url))
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &mqDefault{
		connection: conn,
		channel:    ch,
	}, nil
}

func NewMQFromConnection(conn *amqp.Connection) (MQ, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &mqDefault{
		connection: conn,
		channel:    ch,
	}, nil
}

func (mq *mqDefault) Connection() *amqp.Connection {
	return mq.connection
}

func (mq *mqDefault) Channel() *amqp.Channel {
	return mq.channel
}

func (mq *mqDefault) Queue() amqp.Queue {
	return mq.queue
}

func (mq *mqDefault) QueueDeclare(config *MQConfigQueue) (amqp.Queue, error) {
	q, err := NewQueue(mq.channel, config)
	if err != nil {
		return mq.queue, err
	}

	mq.queue = q
	return mq.queue, nil
}

func (mq *mqDefault) QueueBind(config *MQConfigQueueBind) error {
	err := NewQueueBind(mq.channel, config)
	if err != nil {
		return err
	}

	return nil
}

func (mq *mqDefault) ExchangeDeclare(config *MQConfigExchange) error {
	err := NewExchange(mq.channel, config)
	if err != nil {
		return err
	}

	return nil
}

func (mq *mqDefault) Publish(publish *MQConfigPublish) error {
	return mq.channel.Publish(
		publish.Exchange,
		publish.RoutingKey,
		publish.Mandatory,
		publish.Immediate,
		publish.Message,
	)
}

func (mq *mqDefault) Consume(consume *MQConfigConsume) (<-chan amqp.Delivery, error) {
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

func (mq *mqDefault) Close() {
	mq.channel.Close()
	mq.connection.Close()
}
