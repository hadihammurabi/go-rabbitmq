package gorabbitmq

import "github.com/streadway/amqp"

type MQWithQueue struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      amqp.Queue
}

type MQWithQueueConfig struct {
	Connection *MQConfigConnection
	Queue      *MQConfigQueue
}

func (mq *MQWithQueue) Publish(publish *MQConfigPublish) error {
	return mq.Channel.Publish(
		publish.Exchange,
		publish.RoutingKey,
		publish.Mandatory,
		publish.Immediate,
		publish.Message,
	)
}

func (mq *MQWithQueue) Consume(consume *MQConfigConsume) error {
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
