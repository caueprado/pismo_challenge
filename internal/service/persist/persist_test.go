package persist

import (
	"errors"
	"pismo/internal/domain"
	dbmocks "pismo/internal/infra/db/mocks"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Testando a função Insert com sucesso
func TestPersistenceService_Insert_Success(t *testing.T) {
	// Criar o mock do cliente db.Client
	mockClient := new(dbmocks.Client)

	// Criar o evento de teste
	event := domain.EventMessage{
		UUID:      "test-uuid",
		Context:   "test-context",
		EventType: "test-event",
		Client:    "test-client",
		CreatedAt: time.Now(),
	}

	// Esperar que o método Insert do mock seja chamado com os parâmetros corretos
	mockClient.On("Insert", mock.MatchedBy(func(item map[string]*dynamodb.AttributeValue) bool {
		return *item["UUID"].S == event.UUID &&
			*item["Context"].S == event.Context &&
			*item["EventType"].S == event.EventType &&
			*item["Client"].S == event.Client &&
			*item["CreatedAt"].S == event.CreatedAt.Format(time.RFC3339)
	})).Return(nil)

	// Criar o PersistenceService usando o mock
	service := NewPersistenceService(mockClient)

	// Chamar o método Insert
	err := service.Insert(event)

	// Verificar se não ocorreu erro
	assert.NoError(t, err)

	// Verificar se o mock foi chamado corretamente
	mockClient.AssertExpectations(t)
}

// Testando a função Insert com erro
func TestPersistenceService_Insert_Error(t *testing.T) {
	// Criar o mock do cliente db.Client
	mockClient := new(dbmocks.Client)

	// Criar o evento de teste
	event := domain.EventMessage{
		UUID:      "test-uuid",
		Context:   "test-context",
		EventType: "test-event",
		Client:    "test-client",
		CreatedAt: time.Now(),
	}

	// Esperar que o método Insert do mock retorne um erro
	mockClient.On("Insert", mock.Anything).Return(errors.New("failed to insert item"))

	// Criar o PersistenceService usando o mock
	service := NewPersistenceService(mockClient)

	// Chamar o método Insert
	err := service.Insert(event)

	// Verificar se ocorreu erro
	assert.Error(t, err)
	assert.Equal(t, "failed to insert item", err.Error())

	// Verificar se o mock foi chamado corretamente
	mockClient.AssertExpectations(t)
}
