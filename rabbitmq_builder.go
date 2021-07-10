package gorabbitmq

type MQConfigBuilder struct {
	Connection *MQConfigConnection
	Exchange   *MQConfigExchange
	Queue      *MQConfigQueue
	Bind       *MQConfigQueueBind
}

func NewMQBuilder() *MQConfigBuilder {
	return &MQConfigBuilder{}
}

func (builder *MQConfigBuilder) SetConnection(url string) *MQConfigBuilder {
	builder.Connection = &MQConfigConnection{
		URL: url,
	}
	return builder
}

func (builder *MQConfigBuilder) SetExchange(config *MQConfigExchange) *MQConfigBuilder {
	builder.Exchange = config
	return builder
}

func (builder *MQConfigBuilder) SetQueue(config *MQConfigQueue) *MQConfigBuilder {
	builder.Queue = config
	return builder
}

func (builder *MQConfigBuilder) SetBind(config *MQConfigQueueBind) *MQConfigBuilder {
	builder.Bind = config
	return builder
}

func (builder *MQConfigBuilder) Build() (MQ, error) {
	var mq MQ

	mq, err := NewMQ(builder.Connection.URL)
	if err != nil {
		return nil, err
	}

	if builder.Queue != nil {
		_, err := mq.DeclareQueue(builder.Queue)
		if err != nil {
			return nil, err
		}
	}

	if builder.Exchange != nil {
		err = mq.DeclareExchange(builder.Exchange)
		if err != nil {
			return nil, err
		}

		err = mq.QueueBind(builder.Bind)
		if err != nil {
			return nil, err
		}
	}

	return mq, nil
}
