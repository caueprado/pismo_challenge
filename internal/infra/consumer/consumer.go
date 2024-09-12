package consumer

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	consumerkafka "pismo/internal/infra/kafka"
	"pismo/internal/usecase"
	"syscall"
	"time"

	gokafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

const maxRetries = 5

type EventConsumer interface {
	Consume(doneChan chan bool) error
	consumeMessages(consumer consumerkafka.ConsumerInterface, doneChan chan bool)
}

type eventConsumer struct {
	consumer  consumerkafka.ConsumerInterface
	processor usecase.EventProcessor
}

func NewEventConsumer(
	consumer consumerkafka.ConsumerInterface,
	processor usecase.EventProcessor,
) EventConsumer {
	return &eventConsumer{
		consumer:  consumer,
		processor: processor,
	}
}

func (e *eventConsumer) Consume(doneChan chan bool) error {
	go e.consumeMessages(e.consumer, doneChan)

	// Canal para capturar sinais do sistema
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Bloquear até que receba um sinal de término
	<-sigChan

	// Enviar sinal para parar o consumo
	doneChan <- true

	return nil
}

// Function to consume messages from a Kafka topic
func (e *eventConsumer) consumeMessages(consumer consumerkafka.ConsumerInterface, doneChan chan bool) {
	defer close(doneChan)
	for {
		select {
		case <-doneChan:
			fmt.Println("Stopping message consumption...")
			return
		default:
			msg, err := consumer.ReadMessage()
			if err != nil {
				log.Printf("Error consuming: %v (%v)\n", err, msg)
			} else {
				go func(message []byte) {
					log.Printf("Message consumed from topic %s: %s\n", *msg.TopicPartition.Topic, string(msg.Value))
					e.processWithRetries(msg)
				}(msg.Value)
			}
		}
	}
}

// Função que processa a mensagem com re-tentativas
func (e *eventConsumer) processWithRetries(msg *gokafka.Message) {
	var attempt int
	for attempt = 1; attempt <= maxRetries; attempt++ {
		err := e.processor.ProcessEvent(msg.Value)
		if err != nil {
			log.Printf("Error processing message (attempt %d/%d): %v", attempt, maxRetries, err)
			if attempt < maxRetries {
				// Aguarda um tempo antes de tentar novamente
				time.Sleep(2 * time.Second)
			}
		} else {
			log.Printf("Message processed successfully: %s\n", string(msg.Value))
			return
		}
	}

	// Se falhar após o número máximo de tentativas
	if attempt > maxRetries {
		log.Printf("Message failed after %d attempts: %s\n", maxRetries, string(msg.Value))
		// TODO: enviar a mensagem para um DLQ
	}
}
