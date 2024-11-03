package repository

import (
	"github.com/victornascimento22/api-1.0/internal/database"
	"github.com/victornascimento22/api-1.0/internal/models"
)

func saveUser(user models.User) error {

	db, _ := database.ConnectDatabase()

	defer db.Close()

	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`

	_, err := db.Exec(query, user.Username, user.Email, user.Password)

	if err != nil {
		return err
	}

	return nil
}
