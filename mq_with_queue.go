package gorabbitmq

import (
	"github.com/streadway/amqp"
)

type mqWithQueue struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      amqp.Queue
}

type MQConfigWithQueue struct {
	Connection *MQConfigConnection
	Queue      *MQConfigQueue
}

func NewMQWithQueue(config *MQConfigWithQueue) (MQ, error) {
	mq, err := NewMQ(config.Connection)
	if err != nil {
		return nil, err
	}

	channel := mq.GetChannel()
	q, err := NewQueue(channel, config.Queue)
	if err != nil {
		return nil, err
	}

	return &mqWithQueue{
		Connection: mq.GetConnection(),
		Channel:    channel,
		Queue:      q,
	}, nil
}

func (mq *mqWithQueue) GetConnection() *amqp.Connection {
	return mq.Connection
}

func (mq *mqWithQueue) GetChannel() *amqp.Channel {
	return mq.Channel
}

func (mq *mqWithQueue) GetQueue() amqp.Queue {
	return mq.Queue
}

func (mq *mqWithQueue) Publish(publish *MQConfigPublish) error {
	return mq.Channel.Publish(
		publish.Exchange,
		publish.RoutingKey,
		publish.Mandatory,
		publish.Immediate,
		publish.Message,
	)
}

func (mq *mqWithQueue) Consume(q amqp.Queue, consume *MQConfigConsume) (<-chan amqp.Delivery, error) {
	qname := q.Name
	if consume.Name != "" {
		qname = consume.Name
	}

	consumer, err := mq.Channel.Consume(
		qname,
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
