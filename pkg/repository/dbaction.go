package repository

import (
	"database/sql"
	"fmt"
	"github.com/maxkalayda/lnkCutNew/pkg"
	"log"
)

type Link interface {
	AddLink(ShortLink, OriginalLink string) (int, error)
	SearchRow(sl string) (string, string, error)
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
	query := fmt.Sprintf("INSERT INTO storage_links_tab (short_link, original_link) values ($1, $2)")
	row := d.db.QueryRow(query, ShortLink, OriginalLink)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	log.Println("Row inserted")
	return 0, nil
}

func (d *DBAdd) SearchRow(sl string) (string, string, error) {
	var id int
	//log.Printf("searchrow: %v\n",d.db)
	query := fmt.Sprintf("SELECT * FROM storage_links_tab WHERE short_link = '%s'", sl)
	row, err := d.db.Query(query)
	var l1 string
	var l2 string
	for row.Next() {
		err := row.Scan(&l1, &l2)
		if err == sql.ErrNoRows {
			return sl, "Не существует такой ссылки", err
		}
		log.Println("l1, l2:", l1, ":", l2)
	}
	defer row.Close()
	log.Println("search row:", id, row)
	return l1, l2, err
}

func (m *MapAdd) AddLink() {
	log.Println("short link, origLink:", m.ShortLink, m.OriginalLink)
	pkg.MSync.Store(m.ShortLink, m.OriginalLink)
}
