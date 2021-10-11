package repository

import (
	"database/sql"
	"fmt"
	"github.com/maxkalayda/lnkCutNew/pkg"
	"log"
)

type Link interface {
	AddLink(ShortLink, OriginalLink string) (int, error)
}

//type LinkMap interface {
//	AddLink(ShortLink, OriginalLink string)
//}

type DBAdd struct {
	db           *sql.DB
	ShortLink    string
	OriginalLink string
}

func NewLink(dbf *sql.DB) *DBAdd {
	return &DBAdd{db: dbf}
}

type MapAdd struct {
	ShortLink    string
	OriginalLink string
}

//
//func NewLinkMap (dbf *sync.Map) *MapAdd {
//	return &MapAdd{mp: dbf}
//}

func (d *DBAdd) AddLink(ShortLink, OriginalLink string) (int, error) {
	log.Println("short link, origLink:", ShortLink, OriginalLink)
	var id int
	//db, err := PostgresConnect()
	//if err != nil {
	//	log.Fatalf("failed to init db: %s, %s", err.Error(), db)
	//}
	query := fmt.Sprintf("INSERT INTO storage_links_tab (short_link, original_link) values ($1, $2)")
	row := d.db.QueryRow(query, ShortLink, OriginalLink)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	log.Println("Row inserted")
	return 0, nil
}

func (m *MapAdd) AddLink() {
	log.Println("short link, origLink:", m.ShortLink, m.OriginalLink)
	pkg.MSync.Store(m.ShortLink, m.OriginalLink)
}
