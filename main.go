package main

import (
	"./lib/database"
	"./lib/search"
	"./lib/shaping"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	if os.Args[1] == "kigo_search" {
		srs := search.Search("./kigo.sqlite3", strings.Join(os.Args[2:], " "))
		for _, sr := range srs {
			fmt.Printf("name:%s, class:%s\n", sr.Name, sr.KigoClass)
		}
	} else if os.Args[1] == "shaping" {
		shaping.Shaping(os.Args[2:])
	} else if os.Args[1] == "create_db" {
		database.InsertKigo("./kigo.sqlite3")
	}
}
