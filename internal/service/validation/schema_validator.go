package validation

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

type SchemaValidator interface {
	Validate(data []byte) error
}

type EventValidator struct {
	schemaLoader gojsonschema.JSONLoader
}

func NewEventValidator() (*EventValidator, error) {
	path, err := filepath.Abs(filepath.Join("..", "resource", "schema.json"))
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %v", err)
	}
	urlPath := "file://" + filepath.ToSlash(path)

	// Criar o carregador de esquema
	schemaLoader := gojsonschema.NewReferenceLoader(urlPath)
	return &EventValidator{schemaLoader: schemaLoader}, nil
}

func (e *EventValidator) Validate(data []byte) error {
	documentLoader := gojsonschema.NewStringLoader(string(data))

	// Validar o JSON com o esquema
	result, err := gojsonschema.Validate(e.schemaLoader, documentLoader)
	if err != nil {
		fmt.Printf("Erro ao validar o JSON: %s\n", err)
		return err
	}

	// Verificar o resultado da validação
	if !result.Valid() {
		var sb strings.Builder
		sb.WriteString("JSON does not match schema:\n")
		for _, desc := range result.Errors() {
			sb.WriteString(fmt.Sprintf("- %s\n", desc))
		}
		return errors.New(sb.String())
	}

	return nil
}
