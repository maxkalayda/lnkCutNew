package service

import (
	"database/sql"
	_ "github.com/jmoiron/sqlx"
	"github.com/maxkalayda/lnkCutNew/pkg"
	"log"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/maxkalayda/lnkCutNew/pkg/repository"
)

type Implementation struct {
	LinkRepo repository.Link
}

func NewLinkRepo(db *sql.DB) *Implementation {
	lrep := repository.NewLink(db)
	return &Implementation{LinkRepo: lrep}
}

//type ImplementationMap struct {
//	LinkRepoMap repository.Link
//}
//
//func NewLinkRepoMap (mp *sync.Map) *Implementation {
//	lrep:=repository.NewLinkMap(mp)
//	return &Implementation{LinkRepo:lrep}
//}

func RandomizeString(link string) string {
	alphabet := "a1b2c3d4f5z6x7c8v9b0mnbhj"
	alphabetHelp := "abc1def2ghi3jkl4mnop5qrs6tuv7wxy8zAB9CDEF0GHIJKLMNOPQRSTUVWXYZ5qrs6tuv7wxy8zAB9CDEF0GHEF0GHIJKLMNO4mnop5qrs6tdef2ghi36tuv7wxy8LMNO4mn"
	alphabetDig := "0123456789"
	alphabetLower := "abcdefghijklmnopqrstuvwxyz"
	alphabetUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	originalLink := link
	originalLinkLen := utf8.RuneCountInString(originalLink)
	link = strings.ToLower(link)
	//наличие https/http
	link = strings.ReplaceAll(link, "http://", "")
	link = strings.ReplaceAll(link, "https://", "")
	if utf8.RuneCountInString(link) < 9 {
		link += alphabet[utf8.RuneCountInString(link):9]
	} else if utf8.RuneCountInString(link) > 9 {
		link = link[0:9]
	} else {
		link = link[0:9]
	}
	rLink := []rune(link)
	//преобразуем ссылки
	//здесь баг, если использовать @@, то есть два спец символа подряд и более, то ссылка становится не уникальной
	for i, j := 0, len(rLink)-1; i < j; i, j = i+1, j-1 {
		rLink[i], rLink[j] = rLink[j], rLink[i]
		if rLink[i]%2 == 0 {
			rLink[i] = unicode.ToUpper(rLink[i])
		}
	}
	//преобразование ссылки без спец символов, если в неё поместили спец.символ
	for i, r := range rLink {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			rLink[i] = rune(alphabet[i])
		} else {
			rLink[i] = r
		}
	}
	//для сравнения, что такая линка не юзается
	//pkg.CompareSlice = append(pkg.CompareSlice, string(rLink))
	//проверка перед отправкой дальше, что соответствует шаблону
	//длина 10 символов
	//содержит буквы up down
	//содержит цифры
	//сожержит _
	//уникальна
	if value, ok := pkg.MSync.Load(string(rLink) + "_"); ok {
		log.Println("Ссылка уже существует в MAP!")
		if value != originalLink[0:originalLinkLen] {
			log.Println("Оригинальные ссылки разные", value, originalLink[0:originalLinkLen])
			//пишем логику доработки изменённой ссылки
			tmp := []rune(link[0:1])
			for i, _ := range rLink {
				rLink[i] = rune(alphabetHelp[rune(i)+tmp[0]])
			}

		} else {
			log.Println("Оригинальные ссылки одинаковые", value, originalLink[0:originalLinkLen])
		}

	}
	//проверка на количество апперов и лоуверов
	countDig := 0
	countUpper := 0
	countLower := 0
	for _, r := range rLink {
		if unicode.IsDigit(r) {
			countDig += 1
		} else if unicode.IsUpper(r) {
			countUpper += 1
		} else if unicode.IsLower(r) {
			countLower += 1
		}
	}

	if countDig == 0 {
		rLink[0] = rune(alphabetDig[originalLinkLen%10])
		rLink[8] = rune(alphabetDig[originalLinkLen%5])
	}
	if countUpper == 0 {
		rLink[1] = rune(alphabetUpper[originalLinkLen%10])
		rLink[7] = rune(alphabetUpper[originalLinkLen%5])
	}
	if countLower == 0 {
		rLink[2] = rune(alphabetLower[originalLinkLen%10])
		rLink[6] = rune(alphabetLower[originalLinkLen%5])
	}

	for _, r := range rLink {
		if unicode.IsDigit(r) {
			countDig += 1
		}
	}
	log.Printf("countDig: %d\tcountUpper: %d\tcountLower: %d", countDig, countUpper, countLower)
	log.Println("len url short:", len(rLink), rLink)

	return string(rLink) + "_"
}

func CuttingLink(link string) string {
	//создаём укороченную линку и вносим в мап
	LinkOriginal := link
	link = RandomizeString(link)

	initToMap := repository.MapAdd{ShortLink: link, OriginalLink: LinkOriginal}
	initToMap.AddLink()
	log.Println("Service - addToMap:", initToMap)

	pkg.MSync.Range(func(key, value interface{}) bool {
		log.Println("MSync:", key, value)
		return true
	})

	return link
}
