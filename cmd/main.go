package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {

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
