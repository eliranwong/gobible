// Package bible
// retrieve bible module data
package bible

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"regexp"

	"github.com/eliranwong/gobible/internal/check"
	_ "github.com/mattn/go-sqlite3"
)

func getDb(module string) *sql.DB {
	filePath := fmt.Sprintf("data/bibles/%v.bible", module)
	if !check.FileExists(filePath) {
		filePath = "data/bibles/KJV.bible"
	}
	dbPath := filepath.FromSlash(filePath)
	db, err := sql.Open("sqlite3", dbPath)
	check.DbErr(err)
	return db
}

func SingleVerse(module string, b, c, v int) {
	db := getDb(module)
	defer db.Close()

	if b == 0 {
		b = 1
	}
	if c == 0 {
		c = 1
	}
	query := ""
	if v == 0 {
		query = fmt.Sprintf("SELECT Verse, Scripture from Verses WHERE Book=%v AND Chapter=%v", b, c)
	} else {
		query = fmt.Sprintf("SELECT Verse, Scripture from Verses WHERE Book=%v AND Chapter=%v AND Verse=%v", b, c, v)
	}
	results, err := db.Query(query)
	check.DbErr(err)
	defer results.Close()

	for results.Next() {
		var verseNo int
		var verseText string
		err = results.Scan(&verseNo, &verseText)
		check.DbErr(err)
		verseText = regexp.MustCompile("<[^<>]*?>").ReplaceAllString(verseText, "")
		display := fmt.Sprintf("%v.%v.%v %v", b, c, verseNo, verseText)
		fmt.Println(display)
	}
	err = results.Err()
	check.DbErr(err)
}
