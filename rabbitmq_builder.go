package gorabbitmq

type MQConfigBuilder struct {
	Connection *MQConfigConnection
	Exchange   *MQConfigExchange
	Queue      *MQConfigQueue
	Bind       *MQConfigBind
}

func NewMQBuilder() *MQConfigBuilder {
	return &MQConfigBuilder{}
}

func (builder *MQConfigBuilder) SetConnection(config *MQConfigConnection) *MQConfigBuilder {
	builder.Connection = config
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

func (builder *MQConfigBuilder) SetBind(config *MQConfigBind) *MQConfigBuilder {
	builder.Bind = config
	return builder
}

func (builder *MQConfigBuilder) Build() (MQ, error) {
	var mq MQ

	mq, err := NewMQ(builder.Connection)
	if err != nil {
		return nil, err
	}

	if builder.Queue != nil {
		mq, err = NewMQWithQueue(&MQConfigWithQueue{
			Connection: builder.Connection,
			Queue:      builder.Queue,
		})
		if err != nil {
			return nil, err
		}
	}

	if builder.Exchange != nil {
		mq, err = NewMQWithExchange(&MQWithExchangeConfig{
			Connection: builder.Connection,
			Queue:      builder.Queue,
			Exchange:   builder.Exchange,
			Bind:       builder.Bind,
		})
		if err != nil {
			return nil, err
		}
	}

	return mq, nil
}
