package gorabbitmq

import (
	"github.com/hadihammurabi/go-rabbitmq/connection"
	"github.com/streadway/amqp"
)

const (
	ChannelDefault string = "default"
)

type MQConfig struct {
	Connection *connection.Connection
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
}
