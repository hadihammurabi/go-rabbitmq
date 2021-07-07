
<a name="0.1.1"></a>
## [0.1.1](https://github.com/hadihammurabi/go-rabbitmq/compare/0.1.0...0.1.1) (2021-07-07)

### Chore

* generate change log

### Docs

* how it works
* update mq Consume guide
* update mq Close guide
* update QueueBind with new API
* update NewMQ, DeclareQueue, and DeclareExchange with new API
* consume messages guide
* publish a message guide
* bind queue to exchange guide
* declare queue guide
* connect to rabbitmq guide
* installation guide
* put list of features
* put background
* add license
* contributions guide
* readme content structure

### Feat

* **mq:** can close connection and channel from mq object
* **mq:** can bind queue from mq object
* **mq:** can declare exchange from mq object
* **mq:** can declare queue from mq object

### Fix

* **docs:** code highlight for some examples

### Refactor

* mq Consume queue itself
* new mq only need url


<a name="0.1.0"></a>
## 0.1.0 (2021-07-06)

### Chore

* configure git chglog

### Feat

* **connection:** can connect
* **mq:** builder pattern
* **mq:** can publish and subscribe (fanout)
* **mq:** can publish and subscribe (direct)
* **queue:** can create/declare queue

### Refactor

* manual consumer data listening
* rename MQWithExchangeConfig to MQConfigWithExchange
* rename MQWithQueueConfig to MQConfigWithQueue
* GetQueue contract for mq
* add contract for all mq
* separate new mq and new mq with a queue

