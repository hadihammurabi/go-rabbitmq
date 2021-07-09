package gorabbitmq

import (
	"sync"

	"github.com/streadway/amqp"
)

type mqDefault struct {
	Lock       *sync.Mutex
	Connection *amqp.Connection
	Channel    map[string]*amqp.Channel
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
		Lock:       &sync.Mutex{},
		Channel: map[string]*amqp.Channel{
			"default": ch,
		},
	}, nil
}

func (mq *mqDefault) GetConnection() *amqp.Connection {
	return mq.Connection
}

func (mq *mqDefault) CreateChannel(name string) (*amqp.Channel, error) {

	mq.Lock.Lock()
	ch, err := mq.Connection.Channel()
	if err != nil {
		return nil, err
	}

	mq.Channel[name] = ch
	mq.Lock.Unlock()
	return ch, nil
}

func (mq *mqDefault) GetChannel(name ...string) *amqp.Channel {

	n := "default"

	if len(name) > 0 {
		if name[0] != "" {
			n = name[0]
		}
	}

	return mq.Channel[n]
}

func (mq *mqDefault) CloseChannel(name string) error {

	mq.Lock.Lock()
	if err := mq.Channel[name].Close(); err != nil {
		return err
	}

	delete(mq.Channel, name)
	mq.Lock.Unlock()

	return nil
}

func (mq *mqDefault) GetQueue() amqp.Queue {
	return mq.Queue
}

func (mq *mqDefault) DeclareQueue(name string, config *MQConfigQueue) (amqp.Queue, error) {
	q, err := NewQueue(mq.Channel[name], config)
	if err != nil {
		return mq.Queue, err
	}

	mq.Queue = q
	return mq.Queue, nil
}

func (mq *mqDefault) DeclareExchange(name string, config *MQConfigExchange) error {
	err := NewExchange(mq.Channel[name], config)
	if err != nil {
		return err
	}

	return nil
}

func (mq *mqDefault) QueueBind(name string, config *MQConfigBind) error {
	err := NewQueueBind(mq.Channel[name], config)
	if err != nil {
		return err
	}

	return nil
}

func (mq *mqDefault) Publish(name string, publish *MQConfigPublish) error {
	return mq.Channel[name].Publish(
		publish.Exchange,
		publish.RoutingKey,
		publish.Mandatory,
		publish.Immediate,
		publish.Message,
	)
}

func (mq *mqDefault) Consume(name string, consume *MQConfigConsume) (<-chan amqp.Delivery, error) {
	if consume == nil {
		consume = &MQConfigConsume{}
	}

	consumer, err := mq.Channel[name].Consume(
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
