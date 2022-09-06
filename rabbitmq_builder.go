package gorabbitmq

import (
	queue "github.com/hadihammurabi/go-rabbitmq/queue"
	"github.com/streadway/amqp"
)

type MQConfigBuilder struct {
	ConnectionConfig *MQConfigConnection
	Connection       *amqp.Connection
	ExchangeConfig   *MQConfigExchange
	QueueConfig      *queue.Options
	BindConfig       *MQConfigQueueBind
}

func NewMQBuilder() *MQConfigBuilder {
	return &MQConfigBuilder{}
}

func (builder *MQConfigBuilder) SetConnection(url string) *MQConfigBuilder {
	builder.ConnectionConfig = &MQConfigConnection{
		URL: url,
	}
	return builder
}

func (builder *MQConfigBuilder) WithConnection(conn *amqp.Connection) *MQConfigBuilder {
	builder.Connection = conn
	return builder
}

func (builder *MQConfigBuilder) SetExchange(config *MQConfigExchange) *MQConfigBuilder {
	builder.ExchangeConfig = config
	return builder
}

func (builder *MQConfigBuilder) SetQueue(config *queue.Options) *MQConfigBuilder {
	builder.QueueConfig = config
	return builder
}

func (builder *MQConfigBuilder) SetBind(config *MQConfigQueueBind) *MQConfigBuilder {
	builder.BindConfig = config
	return builder
}

func (builder *MQConfigBuilder) Build() (MQ, error) {
	var mq MQ
	var err error

	if builder.Connection == nil {
		mq, err = New(builder.ConnectionConfig.URL)
		if err != nil {
			return nil, err
		}
	} else {
		mq, err = NewFromConnection(builder.Connection)
		if err != nil {
			return nil, err
		}
	}

	if builder.QueueConfig != nil {
		_, err := mq.QueueDeclare(builder.QueueConfig)
		if err != nil {
			return nil, err
		}
	}

	if builder.ExchangeConfig != nil {
		err = mq.ExchangeDeclare(builder.ExchangeConfig)
		if err != nil {
			return nil, err
		}

		err = mq.QueueBind(builder.BindConfig)
		if err != nil {
			return nil, err
		}
	}

	return mq, nil
}
