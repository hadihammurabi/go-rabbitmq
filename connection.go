package gorabbitmq

import "github.com/streadway/amqp"

type ConfigConnection struct {
	URL string
}

func NewConnection(connection *ConfigConnection) (*amqp.Connection, error) {
	return amqp.Dial(connection.URL)
}
