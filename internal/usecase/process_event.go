package usecase

import (
	"encoding/json"
	"fmt"
	"pismo/internal/domain"
	"pismo/internal/service/persist"
	"pismo/internal/service/sender"
	"pismo/internal/service/validation"
)

type EventProcessor interface {
	ProcessEvent(event []byte) error
}

// EventProcessor é responsável por consumir eventos, validá-los, persistir e enviar para outra integração
type eventProcessor struct {
	Validator          validation.SchemaValidator
	PersistenceService persist.PersistenceService
	SenderService      sender.SenderService
}

// NewEventConsumer cria um novo consumidor com dependências injetadas
func NewEventProcessor(
	validator validation.SchemaValidator,
	persistenceService persist.PersistenceService,
	senderService sender.SenderService,
) EventProcessor {
	return &eventProcessor{
		Validator:          validator,
		PersistenceService: persistenceService,
		SenderService:      senderService,
	}
}

// ProcessEvent processa o evento após consumir do Kafka
func (e *eventProcessor) ProcessEvent(event []byte) error {
	// Valida o evento
	err := e.Validator.Validate(event)
	if err != nil {
		return fmt.Errorf("validation error: %v", err)
	}

	// Persiste o evento
	var eventMessage domain.EventMessage

	// Convertendo o evento []byte para a struct EventMessage
	err = json.Unmarshal(event, &eventMessage)
	if err != nil {
		fmt.Println("Erro ao converter o evento:", err)
		return err
	}

	err = e.PersistenceService.Insert(eventMessage)
	if err != nil {
		return fmt.Errorf("error saving event: %v", err)
	}

	// Envia o evento para outra integração
	err = e.SenderService.SendMessage(eventMessage)
	if err != nil {
		return fmt.Errorf("error sending event: %v", err)
	}

	return nil
}
