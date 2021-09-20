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

func AddNewRow(sl, ol string) (int, error) {
	var id int
	db, err := PostgresConnect()
	if err != nil {
		log.Fatalf("failed to init db: %s, %s", err.Error(), db)
	}
	tx, _ := db.Begin()
	query := fmt.Sprintf("INSERT INTO storage_links_tab (short_link, original_link) values ($1, $2)")
	row := db.QueryRow(query, sl, ol)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	tx.Commit()
	log.Println("Row inserted")
	return 0, nil
}

func SearchRow(sl string) (int, error) {
	var id int
	db, err := PostgresConnect()
	if err != nil {
		log.Fatalf("failed to init db: %s, %s", err.Error(), db)
	}
	tx, _ := db.Begin()
	query := fmt.Sprintf("SELECT * FROM storage_links_tab WHERE short_link like '%$1%'")

	row, _ := db.Query(query, sl)
	for row.Next() {
		var shrtLink string
		var origLink string
		err := row.Scan(shrtLink, origLink)
		if err != nil {
			log.Fatalf("failted to search")
		}
		log.Println("db:", shrtLink, origLink)
	}

	tx.Commit()
	log.Println("search row:", row, id)
	return 0, err
}
