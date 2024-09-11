package consumer

import (
	"encoding/json"
	"errors"
	"pismo/internal/domain"
	"pismo/internal/infra/consumer/mocks"
	usecasemocks "pismo/internal/usecase/mocks"
	"sync"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/stretchr/testify/mock"
)

func TestEventConsumer_ConsumeMessages_Success(t *testing.T) {
	// Criar mocks
	mockKafkaConsumer := new(mocks.ConsumerInterface)
	mockProcessor := new(usecasemocks.EventProcessor)

	// Simular mensagem Kafka
	message, _ := buildKafkaMessage()

	mockKafkaConsumer.EXPECT().
		ReadMessage(mock.Anything).
		Return(message, nil)

	mockProcessor.EXPECT().
		ProcessEvent(mock.Anything).
		Return(nil)

	// Criar um WaitGroup para esperar a conclusão do processamento
	var wg sync.WaitGroup
	wg.Add(1)
	mockProcessor.EXPECT().
		ProcessEvent(mock.AnythingOfType("[]byte")).
		Run(func(args []byte) {
			wg.Done()
		}).Return(nil)

	// Criar instância do EventConsumer
	doneChan := make(chan bool)
	eventConsumer := NewEventConsumer(mockKafkaConsumer, mockProcessor)

	// Iniciar o consumo em uma goroutine
	go eventConsumer.Consume(doneChan)

	// Esperar que a mensagem seja processada
	wg.Wait()

	// Enviar sinal para encerrar o consumo
	doneChan <- true
	close(doneChan) // Fechar o canal para sinalizar que não há mais mensagens

	// Verificar se o processador foi chamado corretamente
	mockProcessor.AssertCalled(t, "ProcessEvent", mock.Anything)
}

func TestEventConsumer_ConsumeMessages_Error(t *testing.T) {
	// Criar mocks
	mockKafkaConsumer := new(mocks.ConsumerInterface)
	mockProcessor := new(usecasemocks.EventProcessor)

	// Configurar mock para erro de leitura
	mockKafkaConsumer.On("ReadMessage", mock.Anything).Return(nil, errors.New("read error"))

	// Criar instância do EventConsumer
	doneChan := make(chan bool)
	eventConsumer := NewEventConsumer(mockKafkaConsumer, mockProcessor)

	go eventConsumer.Consume(doneChan)

	// Enviar sinal para encerrar o consumo
	doneChan <- true

	// Verificar se o processador não foi chamado
	mockProcessor.AssertNotCalled(t, "ProcessEvent", mock.Anything)
}

func buildKafkaMessage() (*kafka.Message, error) {
	// Criar uma instância de EventMessage com todos os campos preenchidos
	eventMessage := domain.EventMessage{
		UUID:      "1234",
		Context:   "test-context",
		EventType: "test-event",
		Client:    "Client_A",
		CreatedAt: time.Now(), // Usar a hora atual para CreatedAt
	}

	// Converter o EventMessage para JSON
	messageValue, err := json.Marshal(eventMessage)
	if err != nil {
		return nil, err
	}

	// Simular mensagem Kafka com todos os dados preenchidos
	return &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     aws.String("test-topic"),
			Partition: 0,
		},
		Value: messageValue,
	}, nil
}
