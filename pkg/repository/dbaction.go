package repository

import (
	"fmt"
	"github.com/maxkalayda/lnkCutNew/pkg"
	"log"
)

type Link interface {
	AddLink()
}

type DBAdd struct {
	ShortLink    string
	OriginalLink string
}

func (d *DBAdd) AddLink() (int, error) {
	log.Println("short link, origLink:", d.ShortLink, d.OriginalLink)
	var id int
	db, err := PostgresConnect()
	if err != nil {
		log.Fatalf("failed to init db: %s, %s", err.Error(), db)
	}
	query := fmt.Sprintf("INSERT INTO storage_links_tab (short_link, original_link) values ($1, $2)")
	row := db.QueryRow(query, d.ShortLink, d.OriginalLink)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	log.Println("Row inserted")
	return 0, nil
}

type MapAdd struct {
	ShortLink    string
	OriginalLink string
}

func (m *MapAdd) AddLink() {
	log.Println("short link, origLink:", m.ShortLink, m.OriginalLink)
	pkg.MSync.Store(m.ShortLink, m.OriginalLink)
}
