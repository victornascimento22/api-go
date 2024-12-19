package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/victornascimento22/api-1.0/internal/models"
	"github.com/victornascimento22/api-1.0/internal/repository"
)

func TestSaveUser(t *testing.T) {
	// Arrange (Preparação)
	user := models.User{
		Username: "João Silva",
		Email:    "joao@exemplo.com",
		Password: "senha123",
	}

	// Act (Ação)
	savedUser, err := repository.SaveUser(user)

	// Assert (Verificação)
	assert.NoError(t, err)
	assert.NotNil(t, savedUser)
	assert.NotEmpty(t, savedUser.ID)
	assert.Equal(t, user.Username, savedUser.Username)
	assert.Equal(t, user.Email, savedUser.Email)
}
