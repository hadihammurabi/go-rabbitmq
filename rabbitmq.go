package gorabbitmq

import (
	"github.com/hadihammurabi/go-rabbitmq/connection"
	"github.com/hadihammurabi/go-rabbitmq/exchange"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	ChannelDefault string = "default"
)

type MQConfig struct {
	Connection *connection.Connection
	Exchange   *exchange.Exchange
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
