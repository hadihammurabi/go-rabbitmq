package connection

import "github.com/streadway/amqp"

type Connection struct {
	URL string

	connection *amqp.Connection
}

func New(url string) (*Connection, error) {
	connData := &Connection{
		URL: url,
	}
	conn, err := connData.dial()
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (c *Connection) Channel() (*amqp.Channel, error) {
	return c.connection.Channel()
}

func (c *Connection) Close() error {
	return c.connection.Close()
}

func (c *Connection) Raw() *amqp.Connection {
	return c.connection
}

func (c *Connection) dial() (*Connection, error) {
	amqpConn, err := amqp.Dial(c.URL)
	if err != nil {
		return nil, err
	}

	conn := &Connection{
		connection: amqpConn,
		URL:        c.URL,
	}

	return conn, nil
}

func From(connection *Connection) *Connection {
	return &Connection{
		connection: connection.connection,
		URL:        connection.URL,
	}
}

func FromAMQP(connection *amqp.Connection) *Connection {
	return &Connection{
		connection: connection,
	}
}
