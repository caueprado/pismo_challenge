package usecase

import (
	"encoding/json"
	"errors"
	"pismo/internal/domain"

	persistmocks "pismo/internal/service/persist/mocks"
	sendermocks "pismo/internal/service/sender/mocks"
	validationmocks "pismo/internal/service/validation/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestEventProcessor_ProcessEvent_ValidationError(t *testing.T) {
	// Mocks das dependências
	mockValidator := new(validationmocks.SchemaValidator)
	mockPersistenceService := new(persistmocks.PersistenceService)
	mockSenderService := new(sendermocks.SenderService)

	// Configurar o mock para retornar erro de validação
	mockValidator.On("Validate", mock.Anything).Return(errors.New("validation failed"))

	// Criar a instância do processor com as dependências mockadas
	processor := NewEventProcessor(mockValidator, mockPersistenceService, mockSenderService)

	// Chamar o método ProcessEvent
	err := processor.ProcessEvent([]byte(`{}`))

	// Verificar se o erro esperado foi retornado
	assert.Error(t, err)
	assert.EqualError(t, err, "validation error: validation failed")

	// Verificar se as dependências corretas foram chamadas
	mockValidator.AssertCalled(t, "Validate", mock.Anything)
	mockPersistenceService.AssertNotCalled(t, "Insert", mock.Anything)
	mockSenderService.AssertNotCalled(t, "SendMessage", mock.Anything)
}

func TestEventProcessor_ProcessEvent_UnmarshalError(t *testing.T) {
	// Mocks das dependências
	mockValidator := new(validationmocks.SchemaValidator)
	mockPersistenceService := new(persistmocks.PersistenceService)
	mockSenderService := new(sendermocks.SenderService)

	// Configurar o mock para validação bem-sucedida
	mockValidator.On("Validate", mock.Anything).Return(nil)

	// Criar a instância do processor com as dependências mockadas
	processor := NewEventProcessor(mockValidator, mockPersistenceService, mockSenderService)

	// Simular evento com JSON inválido
	err := processor.ProcessEvent([]byte(`invalid-json`))

	// Verificar se o erro esperado foi retornado
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid character")

	// Verificar se as dependências corretas foram chamadas
	mockValidator.AssertCalled(t, "Validate", mock.Anything)
	mockPersistenceService.AssertNotCalled(t, "Insert", mock.Anything)
	mockSenderService.AssertNotCalled(t, "SendMessage", mock.Anything)
}

func TestEventProcessor_ProcessEvent_PersistenceError(t *testing.T) {
	// Mocks das dependências
	mockValidator := new(validationmocks.SchemaValidator)
	mockPersistenceService := new(persistmocks.PersistenceService)
	mockSenderService := new(sendermocks.SenderService)

	// Configurar o mock para validação e falha na persistência
	mockValidator.On("Validate", mock.Anything).Return(nil)
	mockPersistenceService.On("Insert", mock.Anything).Return(errors.New("insert failed"))

	// Criar a instância do processor com as dependências mockadas
	processor := NewEventProcessor(mockValidator, mockPersistenceService, mockSenderService)

	// Evento de teste
	event := domain.EventMessage{
		UUID:    "1234",
		Context: "test-context",
		Client:  "Client_A",
	}

	// Chamar o método ProcessEvent
	eventBytes, _ := json.Marshal(event)
	err := processor.ProcessEvent(eventBytes)

	// Verificar se o erro esperado foi retornado
	assert.Error(t, err)
	assert.EqualError(t, err, "error saving event: insert failed")

	// Verificar se as dependências corretas foram chamadas
	mockValidator.AssertCalled(t, "Validate", mock.Anything)
	mockPersistenceService.AssertCalled(t, "Insert", mock.Anything)
	mockSenderService.AssertNotCalled(t, "SendMessage", mock.Anything)
}

func TestEventProcessor_ProcessEvent_SendError(t *testing.T) {
	// Mocks das dependências
	mockValidator := new(validationmocks.SchemaValidator)
	mockPersistenceService := new(persistmocks.PersistenceService)
	mockSenderService := new(sendermocks.SenderService)

	// Configurar o mock para validação e persistência bem-sucedidas, mas falha no envio
	mockValidator.On("Validate", mock.Anything).Return(nil)
	mockPersistenceService.On("Insert", mock.Anything).Return(nil)
	mockSenderService.On("SendMessage", mock.Anything).Return(errors.New("send failed"))

	// Criar a instância do processor com as dependências mockadas
	processor := NewEventProcessor(mockValidator, mockPersistenceService, mockSenderService)

	// Evento de teste
	event := domain.EventMessage{
		UUID:    "1234",
		Context: "test-context",
		Client:  "Client_A",
	}

	// Chamar o método ProcessEvent
	eventBytes, _ := json.Marshal(event)
	err := processor.ProcessEvent(eventBytes)

	// Verificar se o erro esperado foi retornado
	assert.Error(t, err)
	assert.EqualError(t, err, "error sending event: send failed")

	// Verificar se as dependências corretas foram chamadas
	mockValidator.AssertCalled(t, "Validate", mock.Anything)
	mockPersistenceService.AssertCalled(t, "Insert", mock.Anything)
	mockSenderService.AssertCalled(t, "SendMessage", mock.Anything)
}

func TestEventProcessor_ProcessEvent_Success(t *testing.T) {
	// Mocks das dependências
	mockValidator := new(validationmocks.SchemaValidator)
	mockPersistenceService := new(persistmocks.PersistenceService)
	mockSenderService := new(sendermocks.SenderService)

	// Configurar o mock para um processamento bem-sucedido
	mockValidator.On("Validate", mock.Anything).Return(nil)
	mockPersistenceService.On("Insert", mock.Anything).Return(nil)
	mockSenderService.On("SendMessage", mock.Anything).Return(nil)

	// Criar a instância do processor com as dependências mockadas
	processor := NewEventProcessor(mockValidator, mockPersistenceService, mockSenderService)

	// Evento de teste
	event := domain.EventMessage{
		UUID:    "1234",
		Context: "test-context",
		Client:  "Client_A",
	}

	// Chamar o método ProcessEvent
	eventBytes, _ := json.Marshal(event)
	err := processor.ProcessEvent(eventBytes)

	// Verificar que não ocorreu erro
	assert.NoError(t, err)

	// Verificar se as dependências corretas foram chamadas
	mockValidator.AssertCalled(t, "Validate", mock.Anything)
	mockPersistenceService.AssertCalled(t, "Insert", mock.Anything)
	mockSenderService.AssertCalled(t, "SendMessage", mock.Anything)
}
