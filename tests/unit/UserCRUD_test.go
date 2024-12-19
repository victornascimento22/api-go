package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

func TestSaveUser(t *testing.T) {
	// Arrange (Preparação)
	user := User{
		Name:     "João Silva",
		Email:    "joao@exemplo.com",
		Password: "senha123",
	}

	// Act (Ação)
	result, err := SaveUser(user)

	// Assert (Verificação)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.ID)
	assert.Equal(t, user.Name, result.Name)
	assert.Equal(t, user.Email, result.Email)
}
