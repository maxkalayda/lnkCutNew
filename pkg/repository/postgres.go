package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

var (
	shortLink    = "short_link"
	originalLink = "original_link"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func PostgresConnect(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Printf("successfully connected to DB: %s:%s", cfg.Host, cfg.Port)
	return db, nil
}

func AddNewRow(sl, ol string) (int, error) {
	//query := fmt.Sprintf("INSERT INTO storage_links_tab (shortLink, originalLink) values ($1, $2) RETURNING id")
	var id int
	db, err := PostgresConnect(Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("failed to init db: %s, %s", err.Error(), db)
	}
	//query := fmt.Sprintf("select * from storage_links_tab")
	query := fmt.Sprintf("INSERT INTO storage_links_tab (short_link, original_link) values ($1, $2)")
	row := db.QueryRow(query, sl, ol)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return 0, nil
}
