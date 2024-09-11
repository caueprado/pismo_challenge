package consumer

import "github.com/confluentinc/confluent-kafka-go/kafka"

type ConsumerInterface interface {
	ReadMessage(timeoutMs int) (*kafka.Message, error)
	Close() error
}
