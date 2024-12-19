package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victornascimento22/api-1.0/internal/models"
	"github.com/victornascimento22/api-1.0/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func RegisterHandler(c *gin.Context) {

	var input RegisterInput

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	//data binding
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Email = input.Email
	u.Password = string(hashedPassword)

	err = repository.SaveUser(u)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
