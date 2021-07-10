package gorabbitmq

import "github.com/streadway/amqp"

type MQConfigConnection struct {
	URL string
}

func NewConnectionOptions() *MQConfigConnection {
	return &MQConfigConnection{}
}

func (config *MQConfigConnection) SetURL(url string) *MQConfigConnection {
	config.URL = url
	return config
}

func NewConnection(connection *MQConfigConnection) (*amqp.Connection, error) {
	return amqp.Dial(connection.URL)
}
