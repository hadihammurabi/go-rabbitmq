package gorabbitmq

import (
	"github.com/streadway/amqp"
)

type mqWithQueue struct {
	MQ         MQ
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      amqp.Queue
}

type MQConfigWithQueue struct {
	Connection *MQConfigConnection
	Queue      *MQConfigQueue
}

func NewMQWithQueue(config *MQConfigWithQueue) (MQ, error) {
	mq, err := NewMQ(config.Connection.URL)
	if err != nil {
		return nil, err
	}

	channel := mq.GetChannel()
	q, err := NewQueue(channel, config.Queue)
	if err != nil {
		return nil, err
	}

	return &mqWithQueue{
		MQ:         mq,
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

func (mq *mqWithQueue) DeclareQueue(config *MQConfigQueue) (amqp.Queue, error) {
	q, err := mq.MQ.DeclareQueue(config)
	if err != nil {
		return mq.Queue, err
	}

	mq.Queue = q
	return mq.Queue, nil
}

func (mq *mqWithQueue) DeclareExchange(config *MQConfigExchange) error {
	err := NewExchange(mq.Channel, config)
	if err != nil {
		return err
	}

	return nil
}

func (mq *mqWithQueue) QueueBind(config *MQConfigBind) error {
	err := NewQueueBind(mq.Channel, config)
	if err != nil {
		return err
	}

	return nil
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

func (mq *mqWithQueue) Close() {
	mq.Connection.Close()
	mq.Channel.Close()
}
