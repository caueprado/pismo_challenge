package consumer

import (
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"pismo/internal/infra/consumer/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestEventConsumer_Consume(t *testing.T) {
	// Mock do Kafka Consumer gerado pelo Mockery
	mockConsumer := new(mocks.EventConsumer)

	// Canal de finalização para testar a função Consume
	doneChan := make(chan bool)

	// Simulação de comportamento esperado
	mockConsumer.EXPECT().
		Consume(mock.AnythingOfType("chan bool")). // espera receber um canal bool
		Return(nil)
		// Run(func(args mock.Arguments) {
		// 	go func() {
		// 		time.Sleep(1 * time.Second) // Simular um pequeno atraso
		// 		doneChan <- true            // Simular envio do sinal de término
		// 	}()
		// })

	// Inicie a goroutine para capturar o sinal de interrupção
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Testando o método Consume
	go func() {
		err := mockConsumer.Consume(doneChan)
		assert.NoError(t, err)
	}()

	// Simule o envio de um sinal SIGINT para finalizar o consumo
	sigChan <- syscall.SIGINT

	// Verifique se o canal foi fechado corretamente
	select {
	case <-doneChan:
		assert.True(t, true, "Consume finished as expected")
	case <-time.After(2 * time.Second):
		t.Fatal("Test timed out waiting for Consume to finish")
	}
}
