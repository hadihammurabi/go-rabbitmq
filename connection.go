package gorabbitmq

import "github.com/streadway/amqp"

type MQConfigConnection struct {
	URL string
}

func NewConnection(connection *MQConfigConnection) (*amqp.Connection, error) {
	return amqp.Dial(connection.URL)
}
