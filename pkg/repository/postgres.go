package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func PostgresConnect() (*sqlx.DB, error) {
	//лучше конфиг
	Host := "localhost"
	Port := "5432"
	Username := "postgres"
	Password := os.Getenv("DB_PASSWORD")
	DBName := "postgres"
	SSLMode := "disable"
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
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

func SearchRow(sl string) (int, error) {
	var id int
	//items := []*models.Item{}
	//items := models.Item{}
	db, err := PostgresConnect()
	if err != nil {
		log.Fatalf("failed to init db: %s, %s", err.Error(), db)
	}
	query := fmt.Sprintf("SELECT original_link FROM storage_links_tab WHERE short_link=(?)", sl)

	row, _ := db.Query(query)

	log.Println("search row:", id, row)
	return 0, err
}
