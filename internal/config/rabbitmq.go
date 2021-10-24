package config

import (
	"fmt"
	"net/url"
	"os"
)

func NewRabbitMQConfig() string {

	return fmt.Sprintf(
		"amqp://%s:%s@%s:%s%s",
		url.QueryEscape(os.Getenv("RMQ_USER")),
		url.QueryEscape(os.Getenv("RMQ_PASS")),
		os.Getenv("RMQ_HOST"),
		os.Getenv("RMQ_PORT"),
		os.Getenv("RMQ_VHOST"),
	)

}
