package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Função para criar um leitor de configuração simulado com dados mockados
func mockYamlConfigReader() *YamlConfigReader {
	return &YamlConfigReader{
		config: Config{
			Database: struct {
				Name     string "yaml:\"name\""
				Region   string "yaml:\"region\""
				Endpoint string "yaml:\"endpoint\""
				Capacity int    "yaml:\"capacity\""
			}{
				Name:     "TestDB",
				Region:   "us-west-2",
				Endpoint: "http://localhost:8000",
				Capacity: 10,
			},
			Kafka: struct {
				TopicName string "yaml:\"topic-name\""
			}{
				TopicName: "TestTopic",
			},
			AWS: struct {
				secret string "yaml:\"secret\""
				key    string "yaml:\"access-key\""
			}{
				secret: "test-secret",
				key:    "test-access-key",
			},
		},
	}
}

// Teste para o método GetDatabaseName
func TestGetDatabaseName(t *testing.T) {
	reader := mockYamlConfigReader()
	assert.Equal(t, "TestDB", reader.GetDatabaseName())
}

// Teste para o método GetDatabaseRegion
func TestGetDatabaseRegion(t *testing.T) {
	reader := mockYamlConfigReader()
	assert.Equal(t, "us-west-2", reader.GetDatabaseRegion())
}

// Teste para o método GetDatabaseEndpoint
func TestGetDatabaseEndpoint(t *testing.T) {
	reader := mockYamlConfigReader()
	assert.Equal(t, "http://localhost:8000", reader.GetDatabaseEndpoint())
}

// Teste para o método GetDatabaseCapacity
func TestGetDatabaseCapacity(t *testing.T) {
	reader := mockYamlConfigReader()
	assert.Equal(t, 10, reader.GetDatabaseCapacity())
}

// Teste para o método GetTopicName
func TestGetTopicName(t *testing.T) {
	reader := mockYamlConfigReader()
	assert.Equal(t, "TestTopic", reader.GetTopicName())
}

// Teste para o método GetAWSSecret
func TestGetAWSSecret(t *testing.T) {
	reader := mockYamlConfigReader()
	assert.Equal(t, "test-secret", reader.GetAWSSecret())
}

// Teste para o método GetAWSAccessKey
func TestGetAWSAccessKey(t *testing.T) {
	reader := mockYamlConfigReader()
	assert.Equal(t, "test-access-key", reader.GetAWSAccessKey())
}
