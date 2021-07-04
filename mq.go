package gorabbitmq

import "github.com/streadway/amqp"

type MQ struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func (mq *MQ) Publish(publish *MQConfigPublish) error {
	return mq.Channel.Publish(
		publish.Exchange,
		publish.RoutingKey,
		publish.Mandatory,
		publish.Immediate,
		publish.Message,
	)
}

func (mq *MQ) Consume(queue amqp.Queue, consume *MQConfigConsume) error {
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

	go consume.OnMessage(consumer)

	return err
}
