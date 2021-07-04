package gorabbitmq

import "github.com/streadway/amqp"

type rabbitMQ struct {
}

type MQ struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      amqp.Queue
}

type Config struct {
	Connection *ConfigConnection
	Queue      *ConfigQueue
}

type ConfigConsume struct {
	Name      string
	Consumer  string
	AutoACK   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      amqp.Table
	OnMessage func(msgs <-chan amqp.Delivery)
}

type ConfigPublish struct {
	Exchange   string
	RoutingKey string
	Mandatory  bool
	Immediate  bool
	Message    amqp.Publishing
}

func NewRabbitMQ(config *Config) (*MQ, error) {
	conn, err := NewConnection(config.Connection)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := NewQueue(ch, config.Queue)
	if err != nil {
		return nil, err
	}

	return &MQ{
		Connection: conn,
		Channel:    ch,
		Queue:      q,
	}, nil
}

func (mq *MQ) Publish(publish *ConfigPublish) error {
	return mq.Channel.Publish(
		publish.Exchange,
		publish.RoutingKey,
		publish.Mandatory,
		publish.Immediate,
		publish.Message,
	)
}

func (mq *MQ) Consume(consume *ConfigConsume) error {
	qname := mq.Queue.Name
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
