package persist

import (
	"pismo/internal/domain"
	"pismo/internal/infra/db"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type PersistenceService interface {
	Insert(event domain.EventMessage) error
}

type persistenceService struct {
	client db.Client
}

func NewPersistenceService(client db.Client) PersistenceService {
	return &persistenceService{client: client}
}

func (p *persistenceService) Insert(event domain.EventMessage) error {
	item := map[string]*dynamodb.AttributeValue{
		"UUID": {
			S: aws.String(event.UUID),
		},
		"Context": {
			S: aws.String(event.Context),
		},
		"EventType": {
			S: aws.String(event.EventType),
		},
		"Client": {
			S: aws.String(event.Client),
		},
		"CreatedAt": {
			S: aws.String(event.CreatedAt.Format(time.RFC3339)),
		},
	}
	return p.client.Insert(item)
}
