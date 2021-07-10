package gorabbitmq

import "github.com/streadway/amqp"

type MQConfigPublish struct {
	Exchange   string
	RoutingKey string
	Mandatory  bool
	Immediate  bool
	Message    amqp.Publishing
}

func NewPublishOptions() *MQConfigPublish {
	return &MQConfigPublish{}
}

func (config *MQConfigPublish) SetExchange(Exchange string) *MQConfigPublish {
	config.Exchange = Exchange
	return config
}

func (config *MQConfigPublish) SetRoutingKey(RoutingKey string) *MQConfigPublish {
	config.RoutingKey = RoutingKey
	return config
}

func (config *MQConfigPublish) SetMandatory(Mandatory bool) *MQConfigPublish {
	config.Mandatory = Mandatory
	return config
}

func (config *MQConfigPublish) SetImmediate(Immediate bool) *MQConfigPublish {
	config.Immediate = Immediate
	return config
}

func (config *MQConfigPublish) SetMessage(Message amqp.Publishing) *MQConfigPublish {
	config.Message = Message
	return config
}
