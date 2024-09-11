package sender

import (
	"errors"
	"pismo/internal/domain"
	"pismo/internal/infra/sqs/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Teste de sucesso ao enviar mensagem
func TestSenderService_SendMessage_Success(t *testing.T) {
	// Criação do mock do sqsclient.SenderClient
	mockSenderClient := new(mocks.SenderClient)

	// Configurar as filas simuladas
	mockSenderClient.On("GetQueues").Return(map[string]string{
		"Client_A": "https://queue-url/client-a",
	})

	// Configurar o comportamento do envio
	mockSenderClient.On("Send", mock.Anything, mock.AnythingOfType("*sqs.SendMessageInput")).Return(nil)

	// Criar o serviço Sender
	service := NewSenderClient(mockSenderClient)

	// Criar o evento de teste
	event := domain.EventMessage{
		UUID:      uuid.New().String(),
		Context:   "test-context",
		EventType: "test-event",
		Client:    "Client_A",
	}

	// Chamar o método SendMessage
	err := service.SendMessage(event)

	// Verificar se não houve erro
	assert.NoError(t, err)

	// Verificar se o método Send foi chamado corretamente
	mockSenderClient.AssertExpectations(t)
}

// Teste de erro ao enviar mensagem quando não há fila para o cliente
func TestSenderService_SendMessage_NoQueueError(t *testing.T) {
	mockSenderClient := new(mocks.SenderClient)

	// Configurar as filas simuladas sem incluir "Client_B"
	mockSenderClient.On("GetQueues").Return(map[string]string{
		"Client_A": "https://queue-url/client-a",
	})

	// Criar o serviço Sender
	service := NewSenderClient(mockSenderClient)

	// Criar o evento de teste para um cliente sem fila
	event := domain.EventMessage{
		UUID:      uuid.New().String(),
		Context:   "test-context",
		EventType: "test-event",
		Client:    "Client_B", // Não existe fila para "Client_B"
	}

	// Chamar o método SendMessage
	err := service.SendMessage(event)

	// Verificar se o erro esperado ocorreu
	assert.Error(t, err)
	assert.EqualError(t, err, "no queue found for client: Client_B")

	mockSenderClient.AssertNotCalled(t, "Send", mock.Anything, mock.AnythingOfType("*sqs.SendMessageInput"))
}

// Teste de falha no envio para a SQS
func TestSenderService_SendMessage_SendError(t *testing.T) {
	mockSenderClient := new(mocks.SenderClient)

	// Configurar as filas simuladas
	mockSenderClient.On("GetQueues").Return(map[string]string{
		"Client_A": "https://queue-url/client-a",
	})

	// Configurar o comportamento do envio para retornar erro
	mockSenderClient.On("Send", mock.Anything, mock.AnythingOfType("*sqs.SendMessageInput")).Return(errors.New("failed to send message"))

	// Criar o serviço Sender
	service := NewSenderClient(mockSenderClient)

	// Criar o evento de teste
	event := domain.EventMessage{
		UUID:      uuid.New().String(),
		Context:   "test-context",
		EventType: "test-event",
		Client:    "Client_A",
	}

	// Chamar o método SendMessage
	err := service.SendMessage(event)

	// Verificar se o erro esperado ocorreu
	assert.Error(t, err)
	assert.EqualError(t, err, "failed to send message to SQS: failed to send message")

	mockSenderClient.AssertExpectations(t)
}
