package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type Client interface {
	Insert(item map[string]*dynamodb.AttributeValue) error
	// CreateTable(readCapacity, writeCapacity int64) error
}

type DynamoDBClient struct {
	dynamodb             dynamodbiface.DynamoDBAPI
	tableName            string
	keySchema            []*dynamodb.KeySchemaElement
	attributeDefinitions []*dynamodb.AttributeDefinition
}

func NewDynamoDBClient(region, endpoint, tableName string) *DynamoDBClient {
	sess := session.Must(session.NewSession(
		&aws.Config{
			Region:   aws.String(region),
			Endpoint: aws.String(endpoint),
		}))

	svc := dynamodb.New(sess)

	keySchema := []*dynamodb.KeySchemaElement{
		{
			AttributeName: aws.String("UUID"),
			KeyType:       aws.String("HASH"),
		},
	}

	attributeDefinitions := []*dynamodb.AttributeDefinition{
		{
			AttributeName: aws.String("UUID"),
			AttributeType: aws.String("S"),
		},
	}

	return &DynamoDBClient{
		dynamodb:             svc,
		tableName:            tableName,
		keySchema:            keySchema,
		attributeDefinitions: attributeDefinitions,
	}
}

func (c *DynamoDBClient) Insert(item map[string]*dynamodb.AttributeValue) error {
	_, err := c.dynamodb.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(c.tableName),
		Item:      item,
	})
	return err
}
