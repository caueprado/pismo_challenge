package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"pismo/cmd/bootstrap"
	"pismo/internal/infra/configuration"
	"pismo/internal/service/validation"
	"pismo/internal/usecase"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	appConfig, err := configuration.NewYamlConfigReader()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	fmt.Printf("Configuratin database: %s", appConfig.GetDatabaseName())

	// Validator
	validatorService, _ := validation.NewEventValidator()

	// Inicializando DynamoDB client
	persistenceService := bootstrap.CreatePersistenceService(
		appConfig.GetDatabaseRegion(),
		appConfig.GetDatabaseEndpoint(),
		appConfig.GetDatabaseName(),
		int64(appConfig.GetDatabaseCapacity()),
	)

	// sender
	senderService, svcErr := bootstrap.InitializeSenderService()
	if svcErr != nil {
		log.Println("error")
	}

	// usecase
	eventProcessUseCase := usecase.NewEventProcessor(validatorService, persistenceService, senderService)

	// Canal para sinalizar quando o consumo deve ser encerrado
	doneChan := make(chan bool)
	bootstrap.InitializeKafkaConsumer(appConfig.GetTopicName(), doneChan, eventProcessUseCase)
	initializeEngine()

	// Tratamento de sinais de sistema para encerrar o servidor e os serviços corretamente
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs // Aguarda sinal de interrupção
	log.Println("Encerrando aplicação...")

}

func route(r *gin.Engine) {
	// Rota simples de exemplo
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func initializeEngine() {
	// Configurando o Gin
	r := gin.Default()
	route(r)

	// Inicializando o servidor web em uma goroutine
	go func() {
		if err := r.Run(":8081"); err != nil {
			log.Fatalf("Erro ao iniciar servidor web: %v", err)
		}
	}()

	log.Println("Servidor e serviços iniciais foram inicializados.")
}
