package gorabbitmq

import (
	"github.com/streadway/amqp"
)

type mqDefault struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      amqp.Queue
}

func NewMQ(url string) (MQ, error) {
	conn, err := NewConnection(&MQConfigConnection{
		URL: url,
	})
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
	return mq.Queue
}

func (mq *mqDefault) DeclareQueue(config *MQConfigQueue) (amqp.Queue, error) {
	q, err := NewQueue(mq.Channel, config)
	if err != nil {
		return mq.Queue, err
	}

	mq.Queue = q
	return mq.Queue, nil
}

func (mq *mqDefault) DeclareExchange(config *MQConfigExchange) error {
	err := NewExchange(mq.Channel, config)
	if err != nil {
		return err
	}

	return nil
}

func (mq *mqDefault) QueueBind(config *MQConfigBind) error {
	err := NewQueueBind(mq.Channel, config)
	if err != nil {
		return err
	}

	return nil
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

func (mq *mqDefault) Consume(consume *MQConfigConsume) (<-chan amqp.Delivery, error) {
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

func (mq *mqDefault) Close() {
	mq.Connection.Close()
	mq.Channel.Close()
}
