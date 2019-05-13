package database

import (
	"database/sql"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"strings"
	"time"
)

type LinkClass struct {
	Url   string
	Title string
}

func InsertKigo(dbpath string) {
	linkmap := mapLinks()
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	for _, links := range linkmap {
		for _, link := range links {
			time.Sleep(3 * time.Second)
			res, err := http.Get(link.Url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			if res.StatusCode != 200 {
				log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
			}

			doc, err := goquery.NewDocumentFromReader(res.Body)
			if err != nil {
				log.Fatal(err)
			}
			doc.Find("#sites-canvas-main-content table").Slice(3, 4).Each(func(i int, s *goquery.Selection) {
				s.Find("tr").Each(func(j int, s2 *goquery.Selection) {
					list := []string{}
					s2.Find("td").Each(func(k int, s3 *goquery.Selection) {
						list = append(list, strings.Replace(s3.Text(), "\n", ",", -1))
					})
					fmt.Println(list)
					if len(list) >= 4 {
						_, err := db.Exec(
							`INSERT INTO kigo (name, pronunciation, kigo_class, side_topic) VALUES (?, ?, ?, ?)`,
							list[0],
							list[1],
							list[2],
							list[3],
						)
						if err != nil {
							panic(err)
						}
					}
				})
			})
		}
		break
	}
}
func mapLinks() map[string][]LinkClass {
	res, err := http.Get("https://sites.google.com/site/haikukigo/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	linkmap := map[string][]LinkClass{}
	doc.Find(".has-expander .parent div a").Each(func(i int, s *goquery.Selection) {
		t, _ := s.Attr("href")
		if strings.Count(t, "/") >= 4 {
			first_word := string([]rune(s.Text())[0])
			linkmap[first_word] = append(linkmap[first_word],
				LinkClass{Url: "https://sites.google.com" + t, Title: s.Text()})
		}
	})
	return linkmap
}
