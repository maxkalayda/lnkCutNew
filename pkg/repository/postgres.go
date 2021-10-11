package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func PostgresConnect() (*sql.DB, error) {
	//лучше конфиг
	Host := "localhost"
	Port := "5432"
	Username := "postgres"
	Password := os.Getenv("DB_PASSWORD")
	DBName := "postgres"
	SSLMode := "disable"
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		Host, Port, Username, DBName, Password, SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Printf("successfully connected to DB: %s:%s", Host, Port)
	return db, nil
}
