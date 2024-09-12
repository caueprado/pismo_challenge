package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const defaultTimeout = -1

type ConsumerInterface interface {
	ReadMessage() (*kafka.Message, error)
	Close() error
}

type KafkaConsumer struct {
	consumer *kafka.Consumer
}

func NewKafkaConsumer(consumer *kafka.Consumer) ConsumerInterface {
	return KafkaConsumer{consumer: consumer}
}

// Close implements ConsumerInterface.
func (k KafkaConsumer) Close() error {
	return k.consumer.Close()
}

// ReadMessage implements ConsumerInterface.
func (k KafkaConsumer) ReadMessage() (*kafka.Message, error) {
	return k.consumer.ReadMessage(defaultTimeout)
}
