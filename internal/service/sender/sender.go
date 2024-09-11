package sender

import (
	"context"
	"encoding/json"
	"fmt"
	"pismo/internal/domain"
	sqsclient "pismo/internal/infra/sqs"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/uuid"
)

type SenderService interface {
	SendMessage(event domain.EventMessage) error
}

type senderService struct {
	sender sqsclient.SenderClient
}

func NewSenderClient(sender sqsclient.SenderClient) SenderService {
	return &senderService{sender: sender}
}

func (s *senderService) SendMessage(event domain.EventMessage) error {
	// Converte o EventMessage para JSON
	messageBody, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %v", err)
	}

	// Determina a fila com base no campo Client
	queues := s.sender.GetQueues()
	queueURL, ok := queues[event.Client]
	if !ok {
		return fmt.Errorf("no queue found for client: %s", event.Client)
	}

	// Envia a mensagem para a fila SQS
	err = s.sender.Send(context.Background(),
		&sqs.SendMessageInput{
			QueueUrl:               aws.String(queueURL),
			MessageBody:            aws.String(string(messageBody)),
			MessageGroupId:         aws.String("event-group"),
			MessageDeduplicationId: aws.String(uuid.New().String()),
		})
	if err != nil {
		return fmt.Errorf("failed to send message to SQS: %v", err)
	}

	fmt.Printf("Message sent to queue: %s\n", queueURL)
	return nil
}
