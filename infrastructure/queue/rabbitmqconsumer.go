package queue

import (
	"errors"
	"github.com/streadway/amqp"
	"ozon-seller-api/infrastructure/logger"
	"time"
)

type RMQConsumer struct {
	URL   string
	QName string
}

func NewRMQConsumer(url string) (*RMQConsumer, error) {
	if url == "" {
		return nil, errors.New("URL пустой")
	}

	return &RMQConsumer{
		URL: url,
	}, nil
}

func (rmq *RMQConsumer) Consume(qname string) ([]interface{}, error) {

	if qname == "" {
		_err := errors.New("queue name is empty")
		logger.Error(_err)
		return nil, _err
	}

	rmq.QName = qname

	conn, err := amqp.Dial(rmq.URL)
	if err != nil {
		logger.Error("Failed to connect to RabbitMQ")
		return nil, err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		logger.Error("Failed to open a channel")
		return nil, err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		rmq.QName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	var list []interface{}
	for d := range msgs {

		time.Sleep(50 * time.Millisecond)

		list = append(list, d.Body)
		d.Ack(false)
	}

	return list, nil
}
