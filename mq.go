package gorabbitmq

import (
	"github.com/streadway/amqp"
)

type mqDefault struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func NewMQ(config *MQConfigConnection) (MQ, error) {
	conn, err := NewConnection(config)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &mqDefault{
		Connection: conn,
		Channel:    ch,
	}, nil
}

func (mq *mqDefault) GetConnection() *amqp.Connection {
	return mq.Connection
}

func (mq *mqDefault) GetChannel() *amqp.Channel {
	return mq.Channel
}

func (mq *mqDefault) GetQueue() amqp.Queue {
	return amqp.Queue{}
}

func (mq *mqDefault) Publish(publish *MQConfigPublish) error {
	return mq.Channel.Publish(
		publish.Exchange,
		publish.RoutingKey,
		publish.Mandatory,
		publish.Immediate,
		publish.Message,
	)
}

func (mq *mqDefault) Consume(queue amqp.Queue, consume *MQConfigConsume) (<-chan amqp.Delivery, error) {
	qname := queue.Name
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
