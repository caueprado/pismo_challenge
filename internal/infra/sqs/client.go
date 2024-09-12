package sqs

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SenderClient interface {
	Send(ctx context.Context, msg *sqs.SendMessageInput) error
	GetQueues() map[string]string
}

type senderClient struct {
	SQSClient *sqs.Client
	Queues    map[string]string // Mapeamento de Client para URLs de fila
}

func NewSenderClient(ctx context.Context) (SenderClient, error) {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("sa-east-1"),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}

	// Define o mapeamento entre o Client do EventMessage e a URL da fila
	queues := map[string]string{
		"Client_A": "https://sqs.sa-east-1.amazonaws.com/077129836877/Client_A.fifo",
		"Client_B": "https://sqs.sa-east-1.amazonaws.com/077129836877/Client_B.fifo",
		"Client_C": "https://sqs.sa-east-1.amazonaws.com/077129836877/Client_C.fifo",
		"Client_D": "https://sqs.sa-east-1.amazonaws.com/077129836877/Client_D.fifo",
		"Retry":    "https://sqs.sa-east-1.amazonaws.com/077129836877/retry.fifo",
	}

	return &senderClient{
		SQSClient: sqs.NewFromConfig(cfg),
		Queues:    queues,
	}, nil
}

func (s senderClient) Send(ctx context.Context, msg *sqs.SendMessageInput) error {
	_, err := s.SQSClient.SendMessage(ctx, msg)
	if err != nil {
		log.Println(err.Error())
	}

	return nil
}

func (s *senderClient) GetQueues() map[string]string {
	return s.Queues
}
