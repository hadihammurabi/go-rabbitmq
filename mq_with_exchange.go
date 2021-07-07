package gorabbitmq

import (
	"errors"

	"github.com/streadway/amqp"
)

type mqWithExchange struct {
	MQ         MQ
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      amqp.Queue
}

type MQConfigWithExchange struct {
	Connection *MQConfigConnection
	Exchange   *MQConfigExchange
	Queue      *MQConfigQueue
	Bind       *MQConfigBind
}

func NewMQWithExchange(config *MQConfigWithExchange) (MQ, error) {
	if config.Connection == nil {
		return nil, errors.New("connection configuration not given")
	}
	if config.Exchange == nil {
		return nil, errors.New("exchange configuration not given")
	}

	mq, err := NewMQ(config.Connection.URL)
	if err != nil {
		return nil, err
	}

	channel := mq.GetChannel()
	err = NewExchange(channel, config.Exchange)
	if err != nil {
		return nil, err
	}

	var q amqp.Queue
	if config.Queue != nil {
		q, err = NewQueue(channel, config.Queue)
		if err != nil {
			return nil, err
		}

	}

	if config.Bind != nil {
		err = NewQueueBind(channel, config.Bind)
		if err != nil {
			return nil, err
		}
	}

	return &mqWithExchange{
		MQ:         mq,
		Connection: mq.GetConnection(),
		Channel:    channel,
		Queue:      q,
	}, nil
}

func (mq *mqWithExchange) GetConnection() *amqp.Connection {
	return mq.Connection
}

func (mq *mqWithExchange) GetChannel() *amqp.Channel {
	return mq.Channel
}

func (mq *mqWithExchange) GetQueue() amqp.Queue {
	return mq.Queue
}

func (mq *mqWithExchange) DeclareQueue(config *MQConfigQueue) (amqp.Queue, error) {
	q, err := mq.MQ.DeclareQueue(config)
	if err != nil {
		return mq.Queue, err
	}

	mq.Queue = q
	return mq.Queue, nil
}

func (mq *mqWithExchange) DeclareExchange(config *MQConfigExchange) error {
	err := NewExchange(mq.Channel, config)
	if err != nil {
		return err
	}

	return nil
}

func (mq *mqWithExchange) QueueBind(config *MQConfigBind) error {
	err := NewQueueBind(mq.Channel, config)
	if err != nil {
		return err
	}

	return nil
}

func (mq *mqWithExchange) Publish(publish *MQConfigPublish) error {
	return mq.Channel.Publish(
		publish.Exchange,
		publish.RoutingKey,
		publish.Mandatory,
		publish.Immediate,
		publish.Message,
	)
}

func (mq *mqWithExchange) Consume(consume *MQConfigConsume) (<-chan amqp.Delivery, error) {
	if consume == nil {
		consume = &MQConfigConsume{}
	}

	consumer, err := mq.Channel.Consume(
		mq.Queue.Name,
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

func (mq *mqWithExchange) Close() {
	mq.Connection.Close()
	mq.Channel.Close()
}
