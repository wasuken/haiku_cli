package search

import (
	"database/sql"
	"fmt"
	"github.com/bluele/mecab-golang"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strings"
)

type SearchResult struct {
	Name          string
	Pronunciation string
	KigoClass     string
	SideTopic     string
}

func listNoum(text string) []string {
	m, err := mecab.New("-Owakati")
	if err != nil {
		fmt.Printf("Mecab instance error. err: %v", err)
	}
	defer m.Destroy()
	tg, err := m.NewTagger()
	if err != nil {
		panic(err)
	}
	defer tg.Destroy()
	lt, err := m.NewLattice(text)
	if err != nil {
		panic(err)
	}
	defer lt.Destroy()

	node := tg.ParseToNode(lt)
	var list []string
	for {
		features := strings.Split(node.Feature(), ",")
		if features[0] == "名詞" {
			list = append(list, node.Surface())
		}
		if node.Next() != nil {
			break
		}
	}
	return list
}
func Search(dbpath, text string) []SearchResult {
	noums := listNoum(text)
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var list []SearchResult
	for _, noum := range noums {
		rows, err := db.Query(
			`select name, pronunciation, kigo_class, side_topic from kigo where name like ?`,
			"%"+noum+"%",
		)
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			var name, pronunciation, kigoClass, sideTopic string
			if err := rows.Scan(&name, &pronunciation, &kigoClass, &sideTopic); err != nil {
				log.Fatal("rows.Scan()", err)
			} else {
				list = append(list, SearchResult{Name: name,
					Pronunciation: pronunciation, KigoClass: kigoClass,
					SideTopic: sideTopic})
			}
		}
	}
	return list
}
