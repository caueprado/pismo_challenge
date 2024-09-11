package validation

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Teste para a função NewEventValidator
func TestNewEventValidator(t *testing.T) {
	validator, err := NewEventValidator()

	// Verifica se o validator foi criado corretamente
	assert.NoError(t, err)
	assert.NotNil(t, validator)
	assert.NotNil(t, validator.schemaLoader)
}

// Teste para o método Validate quando o JSON é válido
func TestEventValidator_Validate_ValidJSON(t *testing.T) {
	validator, err := NewEventValidator()
	assert.NoError(t, err)

	// JSON de exemplo válido
	validJSON := `{"key": "value"}`

	// Mockando a função Validate

	// Verificando a validação de um JSON válido
	err = validator.Validate([]byte(validJSON))
	assert.NoError(t, err)
}

// Teste para o método Validate quando o JSON é inválido
func TestEventValidator_Validate_InvalidJSON(t *testing.T) {
	validator, err := NewEventValidator()
	assert.NoError(t, err)

	// JSON de exemplo inválido
	invalidJSON := `{"invalid_key": "value"}`

	// Simulando um erro na validação do JSON
	err = validator.Validate([]byte(invalidJSON))
	assert.Error(t, err)

	// Verificando a mensagem de erro retornada
	expectedError := "JSON does not match schema:"
	assert.True(t, strings.Contains(err.Error(), expectedError))
}
