package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDatabase() error {

	var (
		DbHost     = os.Getenv("DB_HOST")
		DbPort     = os.Getenv("DB_PORT")
		DbUser     = os.Getenv("DB_USER")
		DbPassword = os.Getenv("DB_PASSWORD")
		DbName     = os.Getenv("DB_NAME")
	)

	err := godotenv.Load(".env")

	if err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		DbHost, DbPort, DbUser, DbPassword, DbName)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	err = db.Ping()

	if err != nil {
		return fmt.Errorf("error pinging database: %v", err)
	}

	DB = db
	fmt.Println("Successfully connected!")
	return nil
}
