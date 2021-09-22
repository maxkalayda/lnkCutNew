package repository

import (
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

func (d *DBAdd) AddLink() {
	log.Println("short link, origLink:", d.ShortLink, d.OriginalLink)

}

type MapAdd struct {
	ShortLink    string
	OriginalLink string
}

func (m *MapAdd) AddLink() {
	log.Println("short link, origLink:", m.ShortLink, m.OriginalLink)
	pkg.MSync.Store(m.ShortLink, m.OriginalLink)
}
