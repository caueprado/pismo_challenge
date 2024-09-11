package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"

	"pismo/internal/domain"

	"github.com/google/uuid"
)

// Função para gerar uma nova mensagem de evento
func generateRandomEvent() domain.EventMessage {
	// Gerar um UID único para a mensagem
	uid := uuid.NewString()

	// Contextos possíveis
	contexts := []string{
		"monitoramento", "aplicação de usuários", "autorizador de transações", "integrações",
	}

	// Tipos de evento possíveis
	eventTypes := []string{
		"INSERT", "UPDATE", "DELETE",
	}

	// Clientes possíveis
	clients := []string{
		"Client_A", "Client_B", "Client_C", "Client_D",
	}

	// Selecionar valores aleatórios
	context := contexts[rand.Intn(len(contexts))]
	eventType := eventTypes[rand.Intn(len(eventTypes))]
	client := clients[rand.Intn(len(clients))]

	// Criar a mensagem de evento
	return domain.EventMessage{
		UUID:      uid,
		Context:   context,
		EventType: eventType,
		Client:    client,
		CreatedAt: time.Now(),
	}
}

// Função para escrever a mensagem JSON no arquivo em formato NDJSON
func writeEventToFile(event domain.EventMessage, file *os.File) error {
	// Transformar a mensagem em JSON
	jsonData, err := json.Marshal(event)
	if err != nil {
		return err
	}

	// Adicionar uma nova linha após cada objeto JSON (NDJSON)
	_, err = file.WriteString(string(jsonData) + "\n")
	if err != nil {
		return err
	}

	return nil
}

func execute() {
	// Semente do gerador de números aleatórios
	rand.Seed(time.Now().UnixNano())

	// Abrir o arquivo em modo de append (cria o arquivo se não existir)
	file, err := os.OpenFile("events1.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v", err)
	}
	defer file.Close()

	// Gerar eventos a cada 5 segundos
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Gerar um evento aleatório
			event := generateRandomEvent()

			// Exibir a mensagem no console
			log.Printf("Gerando mensagem: %+v\n", event)

			// Escrever a mensagem no arquivo
			err := writeEventToFile(event, file)
			if err != nil {
				log.Printf("Erro ao escrever mensagem no arquivo: %v\n", err)
			}
		}
	}
}

func main() {
	// Iniciar a geração de eventos
	execute()
}
