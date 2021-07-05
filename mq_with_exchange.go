package gorabbitmq

import (
	"errors"

	"github.com/streadway/amqp"
)

type mqWithExchange struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      amqp.Queue
}

type MQWithExchangeConfig struct {
	Connection *MQConfigConnection
	Exchange   *MQConfigExchange
	Queue      *MQConfigQueue
	Bind       *MQConfigBind
}

func NewMQWithExchange(config *MQWithExchangeConfig) (MQ, error) {
	if config.Connection == nil {
		return nil, errors.New("connection configuration not given")
	}
	if config.Exchange == nil {
		return nil, errors.New("exchange configuration not given")
	}

	mq, err := NewMQ(config.Connection)
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

func (mq *mqWithExchange) Publish(publish *MQConfigPublish) error {
	return mq.Channel.Publish(
		publish.Exchange,
		publish.RoutingKey,
		publish.Mandatory,
		publish.Immediate,
		publish.Message,
	)
}

func (mq *mqWithExchange) Consume(q amqp.Queue, consume *MQConfigConsume) error {
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

	go consume.OnMessage(consumer)

	return err
}
