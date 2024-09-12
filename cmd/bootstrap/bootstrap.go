package bootstrap

import (
	"context"
	"fmt"
	"log"
	"pismo/internal/infra/consumer"
	"pismo/internal/infra/db"
	"pismo/internal/infra/kafka"
	"pismo/internal/infra/sqs"
	"pismo/internal/service/persist"
	"pismo/internal/service/sender"
	"pismo/internal/usecase"

	consumerkafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// Methods to initialize kafka consumer
func NewEventConsumer(topic string, processor usecase.EventProcessor) consumer.EventConsumer {
	eventConsumer := newConsumer(topic)
	kafkaConsumer := kafka.NewKafkaConsumer(eventConsumer)
	return consumer.NewEventConsumer(kafkaConsumer, processor)
}

func newConsumer(topic string) *consumerkafka.Consumer {
	// Configure the Kafka consumer
	consumer, err := consumerkafka.NewConsumer(&consumerkafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "go-consumer-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to create consumer: %s\n", err)
	}
	// Subscribe to the topic
	err = consumer.Subscribe(topic, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %s\n", err)
	}
	return consumer
}

func InitializeKafkaConsumer(
	topic string,
	doneChan chan bool,
	processor usecase.EventProcessor,
) {
	kafkaConsumer := NewEventConsumer(topic, processor)
	// Iniciar o Kafka consumer em uma goroutine
	go func() {
		if err := kafkaConsumer.Consume(doneChan); err != nil {
			log.Fatalf("Erro ao consumir mensagens: %v", err)
		}
	}()
}

// Methods to initialize DB and persistence service
func CreatePersistenceService(
	region,
	endpoint,
	tableName string,
	capacity int64,
) persist.PersistenceService {
	client := db.NewDynamoDBClient(region, endpoint, tableName)

	return persist.NewPersistenceService(client)
}

// Methods to initialize SQS and sender service
func InitializeSenderService() (sender.SenderService, error) {
	ctx := context.Background()

	// Cria uma nova instância do SQSClient
	client, err := sqs.NewSenderClient(ctx)
	if err != nil {
		fmt.Println("Erro ao criar o cliente SQS:", err)
		return nil, err
	}

	senderService := sender.NewSenderClient(client)

	return senderService, nil
}
