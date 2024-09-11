package configuration

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Estrutura que representa o conteúdo do YAML
type Config struct {
	Database struct {
		Name     string `yaml:"name"`
		Region   string `yaml:"region"`
		Endpoint string `yaml:"endpoint"`
		Capacity int    `yaml:"capacity"`
	} `yaml:"database"`
	Kafka struct {
		TopicName string `yaml:"topic-name"`
	} `yaml:"kafka"`
	AWS struct {
		secret string `yaml:"secret"`
		key    string `yaml:"access-key"`
	} `yaml:"aws"`
}

// Interface de configuração
type ConfigReader interface {
	GetDatabaseName() string
	GetDatabaseRegion() string
	GetDatabaseEndpoint() string
	GetDatabaseCapacity() int

	GetTopicName() string

	GetAWSSecret() string
	GetAWSAccessKey() string
}

// Estrutura que implementa a interface
type YamlConfigReader struct {
	config Config
}

// Função para ler o arquivo YAML e retornar o ConfigReader
func NewYamlConfigReader() (ConfigReader, error) {
	configPath := filepath.Join("..", "configs", "properties.yaml")

	// Ler o arquivo YAML
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	// Deserializar o YAML em uma estrutura Config
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal yaml: %v", err)
	}

	return &YamlConfigReader{config: config}, nil
}

// Implementação do método para obter o nome do banco de dados
func (y *YamlConfigReader) GetDatabaseName() string {
	return y.config.Database.Name
}

// Implementação do método para obter capacidade do banco de dados
func (y *YamlConfigReader) GetDatabaseCapacity() int {
	return y.config.Database.Capacity
}

// Implementação do método para obter o nome do tópico
func (y *YamlConfigReader) GetTopicName() string {
	return y.config.Kafka.TopicName
}

// Implementação do método para obter capacidade do banco de dados
func (y *YamlConfigReader) GetDatabaseRegion() string {
	return y.config.Database.Region
}

// Implementação do método para obter o nome do tópico
func (y *YamlConfigReader) GetDatabaseEndpoint() string {
	return y.config.Database.Endpoint
}

// Implementação do método para secret
func (y *YamlConfigReader) GetAWSSecret() string {
	return y.config.AWS.secret
}

// Implementação do método access key
func (y *YamlConfigReader) GetAWSAccessKey() string {
	return y.config.AWS.key
}
