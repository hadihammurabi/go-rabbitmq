package gorabbitmq

import (
	"fmt"
	"sync"

	"github.com/streadway/amqp"
)

type mqDefault struct {
	Lock       *sync.Mutex
	Connection *amqp.Connection
	channel    map[string]*amqp.Channel
	Queue      amqp.Queue
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
		Connection: conn,
		Lock:       &sync.Mutex{},
		channel: map[string]*amqp.Channel{
			"default": ch,
		},
	}, nil
}

func (mq *mqDefault) GetConnection() *amqp.Connection {
	return mq.Connection
}

func (mq *mqDefault) CreateChannel(name ...string) (*amqp.Channel, error) {
	n := ChannelDefault

	if len(name) > 0 {
		if name[0] != "" {
			n = name[0]
		}
	}

	mq.Lock.Lock()
	ch, err := mq.Connection.Channel()
	if err != nil {
		return nil, err
	}

	mq.channel[n] = ch
	mq.Lock.Unlock()
	return ch, nil
}

func (mq *mqDefault) GetChannel(name ...string) *amqp.Channel {
	n := ChannelDefault

	if len(name) > 0 {
		if name[0] != "" {
			n = name[0]
		}
	}

	return mq.channel[n]
}

func (mq *mqDefault) WithChannel(name ...string) MQ {
	fmt.Println("mantap")
	n := ChannelDefault

	if len(name) > 0 {
		if name[0] != "" {
			n = name[0]
		}
	}

	c := mq.GetChannel(n)

	return &mqDefault{
		Connection: mq.Connection,
		Lock:       &sync.Mutex{},
		channel: map[string]*amqp.Channel{
			"default": c,
		},
	}
}

func (mq *mqDefault) GetQueue() amqp.Queue {
	return mq.Queue
}

func (mq *mqDefault) DeclareQueue(config *MQConfigQueue) (amqp.Queue, error) {
	q, err := NewQueue(mq.GetChannel(), config)
	if err != nil {
		return mq.Queue, err
	}

	mq.Queue = q
	return mq.Queue, nil
}

func (mq *mqDefault) DeclareExchange(config *MQConfigExchange) error {
	err := NewExchange(mq.GetChannel(), config)
	if err != nil {
		return err
	}

	return nil
}

func (mq *mqDefault) QueueBind(config *MQConfigBind) error {
	err := NewQueueBind(mq.GetChannel(), config)
	if err != nil {
		return err
	}

	return nil
}

func (mq *mqDefault) Publish(publish *MQConfigPublish) error {
	return mq.GetChannel().Publish(
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

	consumer, err := mq.GetChannel().Consume(
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
}
