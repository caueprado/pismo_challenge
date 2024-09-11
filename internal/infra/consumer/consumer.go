package consumer

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"pismo/internal/usecase"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const maxRetries = 5

type EventConsumer interface {
	Consume(doneChan chan bool) error
	consumeMessages(consumer *kafka.Consumer, doneChan chan bool)
}

type eventConsumer struct {
	consumer  *kafka.Consumer
	processor usecase.EventProcessor
}

func NewEventConsumer(
	consumer *kafka.Consumer,
	processor usecase.EventProcessor,
) EventConsumer {
	return &eventConsumer{
		consumer:  consumer,
		processor: processor,
	}
}

// Function to consume messages from a Kafka topic
func (e *eventConsumer) consumeMessages(consumer *kafka.Consumer, doneChan chan bool) {
	defer close(doneChan) // Fechar o canal quando a goroutine terminar

	for {
		select {
		case <-doneChan:
			// Recebeu o sinal para parar o consumo
			fmt.Println("Stopping message consumption...")
			return
		default:
			msg, err := consumer.ReadMessage(-1) // Timeout: -1 waits indefinitely
			if err != nil {
				log.Printf("Error consuming: %v (%v)\n", err, msg)
			} else {
				// Iniciar o processamento do serviço em uma goroutine
				go func(message []byte) {
					log.Printf("Message consumed from topic %s: %s\n", *msg.TopicPartition.Topic, string(msg.Value))
					e.processWithRetries(msg)
				}(msg.Value)
			}
		}
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

// Função que processa a mensagem com re-tentativas
func (e *eventConsumer) processWithRetries(msg *kafka.Message) {
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
