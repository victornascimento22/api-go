package repository

import (
	"fmt"

	"github.com/victornascimento22/api-1.0/internal/database"
	"github.com/victornascimento22/api-1.0/internal/models"
)

func SaveUser(user models.User) (models.User, error) {
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return models.User{}, fmt.Errorf("campos obrigatórios não podem estar vazios")
	}

	db, err := database.ConnectDatabase()
	if err != nil {
		return models.User{}, fmt.Errorf("erro na conexão com o banco de dados: %w", err)
	}
	defer db.Close()

	query := `INSERT INTO users (username, email, password) 
			  VALUES ($1, $2, $3) RETURNING id`

	var id int
	err = db.QueryRow(query, user.Username, user.Email, user.Password).Scan(&id)
	if err != nil {
		return models.User{}, fmt.Errorf("erro ao salvar usuário: %w", err)
	}

	return user, nil
}

func GetUserbyID(id int) (models.User, error) {
	db, err := database.ConnectDatabase()

	if err != nil {
		fmt.Printf("Erro na conexão com o banco de dados: %v\n", err)
		return models.User{}, err
	}

	defer db.Close()

	query := `SELECT * FROM users WHERE id = $1`

	var user models.User

	err = db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)

	if err != nil {
		return models.User{}, err
	} else {
		return user, nil
	}
}

func GetUserByEmail(email string) (models.User, error) {
	db, err := database.ConnectDatabase()
	if err != nil {
		return models.User{}, fmt.Errorf("erro na conexão com o banco de dados: %w", err)
	}
	defer db.Close()

	query := `SELECT id, username, email, password FROM users WHERE email = $1`

	var user models.User
	err = db.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return models.User{}, fmt.Errorf("usuário não encontrado: %w", err)
	}

	return user, nil
}
