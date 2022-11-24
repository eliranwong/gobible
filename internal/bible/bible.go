// Package bible
// retrieve bible module data
package bible

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"regexp"

	"github.com/eliranwong/gobible/internal/check"
	"github.com/eliranwong/gobible/internal/parser"
	"github.com/eliranwong/gobible/internal/shortcuts"
	_ "github.com/mattn/go-sqlite3"
)

var Default string = "NET"

func getDb(module string) *sql.DB {
	filePath := fmt.Sprintf("data/bibles/%v.bible", module)
	if check.FileExists(filePath) {
		Default = module
	} else {
		filePath = fmt.Sprintf("data/bibles/%v.bible", Default)
	}
	dbPath := filepath.FromSlash(filePath)
	db, err := sql.Open("sqlite3", dbPath)
	check.DbErr(err)
	return db
}

// read bible reference
func Read(module string, references [][]int) {
	shortcuts.Divider()
	if !(len(references) > 0) {
		//skip
	} else if len(references) == 1 && len(references[0]) == 3 {
		ReadChapter(module, references[0])
	} else {
		for _, bcv := range references {
			if len(bcv) > 3 {
				ReadMultiple(module, bcv)
			} else {
				ReadSingle(module, bcv)
			}
		}
	}
}

func ReadChapter(module string, bcv []int) {
	ReadSingle(module, []int{bcv[0], bcv[1], 0})
	shortcuts.Divider()
	ReadSingle(module, bcv)
}

// read either single verse or single chapter
func ReadSingle(module string, bcv []int) {
	db := getDb(module)
	defer db.Close()

	var b, c, v int
	b = bcv[0]
	c = bcv[1]
	v = bcv[2]

	if b == 0 {
		b = 1
	}
	if c == 0 {
		c = 1
	}
	query := ""
	if v == 0 {
		query = fmt.Sprintf("SELECT * from Verses WHERE Book=%v AND Chapter=%v", b, c)
	} else {
		query = fmt.Sprintf("SELECT * from Verses WHERE Book=%v AND Chapter=%v AND Verse=%v", b, c, v)
	}
	results, err := db.Query(query)
	check.DbErr(err)
	processResults(results, err)
}

// read multiple verses
func ReadMultiple(module string, bcv []int) {
	var b, c, cs, vs, ce, ve int
	b = bcv[0]
	cs = bcv[1]
	vs = bcv[2]

	if len(bcv) == 4 {
		ve = bcv[3]
		displayChapterVerses(module, b, cs, vs, ve)
	} else if len(bcv) == 5 {
		ce = bcv[3]
		ve = bcv[4]

		if cs > ce {
			// skip
		} else if cs == ce {
			displayChapterVerses(module, b, cs, vs, ve)
		} else {
			c = cs
			displayChapterVerses(module, b, c, vs, 0)
			c += 1
			for c < ce {
				displayChapterVerses(module, b, c, 0, 0)
				c += 1
			}
			if c == ce {
				displayChapterVerses(module, b, c, 1, ve)
			}
		}
	}
}

// read chapter verses
func displayChapterVerses(module string, b, c, vs, ve int) {
	db := getDb(module)
	defer db.Close()

	var query string
	if vs <= 0 && ve <= 0 {
		query = fmt.Sprintf("SELECT DISTINCT * FROM Verses WHERE Book=%v AND Chapter=%v ORDER BY Verse", b, c)
	} else if vs > 0 && ve > 0 {
		query = fmt.Sprintf("SELECT DISTINCT * FROM Verses WHERE Book=%v AND Chapter=%v AND Verse BETWEEN %v AND %v ORDER BY Verse", b, c, vs, ve)
	} else if vs > 0 {
		query = fmt.Sprintf("SELECT DISTINCT * FROM Verses WHERE Book=%v AND Chapter=%v AND Verse>=%v ORDER BY Verse", b, c, vs)
	}

	results, err := db.Query(query)
	check.DbErr(err)
	processResults(results, err)
}

// process sqlite query results
func processResults(results *sql.Rows, err error) {
	defer results.Close()

	for results.Next() {
		var b, c, v int
		var text string
		err = results.Scan(&b, &c, &v, &text)
		check.DbErr(err)
		text = regexp.MustCompile("<[^<>]*?>").ReplaceAllString(text, "")
		display := fmt.Sprintf("%v %v", parser.BcvToVerseReference([]int{b, c, v}), text)
		displayResults(display)
	}
	err = results.Err()
	check.DbErr(err)
}

func displayResults(text string) {
	fmt.Println(text)
}
