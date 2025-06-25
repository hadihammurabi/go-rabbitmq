package gorabbitmq

import (
	"github.com/hadihammurabi/go-rabbitmq/connection"
	"github.com/rabbitmq/amqp091-go"
)

type NewOptions struct {
	url        string
	connection *connection.Connection
	amqp       *amqp091.Connection
}

func (opt *NewOptions) Apply(opts ...func(*NewOptions)) {
	for _, o := range opts {
		o(opt)
	}
}

func WithURL(url string) func(*NewOptions) {
	return func(opts *NewOptions) {
		opts.url = url
	}
}

func WithConnection(conn *connection.Connection) func(*NewOptions) {
	return func(opts *NewOptions) {
		opts.connection = conn
	}
}

func WithAMQP(conn *amqp091.Connection) func(*NewOptions) {
	return func(opts *NewOptions) {
		opts.amqp = conn
	}
}
