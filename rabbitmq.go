package gorabbitmq

import "github.com/streadway/amqp"

type MQ interface {
	GetConnection() *amqp.Connection
	GetChannel() *amqp.Channel
	Publish(publish *MQConfigPublish) error
	Consume(queue amqp.Queue, consume *MQConfigConsume) error
}

type MQConfig struct {
	Connection *MQConfigConnection
	Exchange   *MQConfigExchange
}

type MQConfigConsume struct {
	Name      string
	Consumer  string
	AutoACK   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      amqp.Table
	OnMessage func(msgs <-chan amqp.Delivery)
}

type MQConfigPublish struct {
	Exchange   string
	RoutingKey string
	Mandatory  bool
	Immediate  bool
	Message    amqp.Publishing
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

// func NewMQWithExchange(config *MQWithExchangeConfig) (*MQWithExchange, error) {
// 	mq, err := NewMQ(config.Connection)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = NewExchange(mq.Channel, config.Exchange)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var q amqp.Queue
// 	if config.Queue != nil {
// 		_, err := NewQueue(mq.Channel, config.Queue)
// 		if err != nil {
// 			return nil, err
// 		}

// 		err = NewQueueBind(mq.Channel, config.Bind)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	return &MQWithExchange{
// 		Connection: mq.Connection,
// 		Channel:    mq.Channel,
// 		Queue:      q,
// 	}, nil
// }
