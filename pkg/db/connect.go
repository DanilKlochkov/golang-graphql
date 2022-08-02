package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Init() (*sql.DB, error) {
	con, err := getEnv()
	if err != nil {
		return nil, err
	}

	db, err := connect(con)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getEnv() (string, error) {
	err := godotenv.Load("C:\\Users\\danii\\golang-graphql\\configs\\variables.env")
	if err != nil {
		return "", err
	}

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("USER"),
		os.Getenv("PASS"),
		os.Getenv("NAME"),
	)

	return connectionString, err
}

func connect(con string) (*sql.DB, error) {
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}
	return db, nil
}
