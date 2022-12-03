package bible

import (
	"database/sql"
	"fmt"
	"path/filepath"

	"github.com/eliranwong/gobible/internal/check"
	"github.com/eliranwong/gobible/internal/parser"
	"github.com/eliranwong/gobible/internal/share"
)

// query = "SELECT Information FROM ScrollMapper WHERE Book=? AND Chapter=? AND Verse=?"
func getXrefDb() *sql.DB {
	dbPath := filepath.Join(share.Data, filepath.FromSlash("xref/cross-reference.sqlite"))
	//dbPath := filepath.FromSlash(filePath)
	db, err := sql.Open("sqlite3", dbPath)
	check.DbErr(err)
	return db
}

func GetXrefs(reference string) [][]string {
	var resultsSlice [][]string = [][]string{}
	// parse bible reference
	references := parser.ExtractAllReferences(reference, false)
	if len(references) > 0 {
		b, c, v := references[0][0], references[0][1], references[0][2]
		db := getXrefDb()
		defer db.Close()
		query := fmt.Sprintf("SELECT Information FROM ScrollMapper WHERE Book=%v AND Chapter=%v AND Verse=%v", b, c, v)
		results, err := db.Query(query)
		check.DbErr(err)

		var information string
		for results.Next() {
			err = results.Scan(&information)
			check.DbErr(err)
		}
		err = results.Err()
		check.DbErr(err)

		information = fmt.Sprintf("%v; %v", reference, information)
		xrefs := parser.ExtractAllReferences(information, false)
		var resultSlice []string

		for _, xref := range xrefs {
			b, c, v := xref[0], xref[1], xref[2]
			query := fmt.Sprintf("SELECT * from Verses WHERE Book=%v AND Chapter=%v AND Verse=%v", b, c, v)
			for _, module := range share.SelectedBibles {
				db := getDb(module)
				defer db.Close()
				results, err := db.Query(query)
				check.DbErr(err)

				for results.Next() {
					var b, c, v int
					var text string
					err = results.Scan(&b, &c, &v, &text)
					check.DbErr(err)
					text = formatVerseText(text)
					resultSlice = []string{module, parser.BcvToVerseReference([]int{b, c, v}), text}
					resultsSlice = append(resultsSlice, resultSlice)
				}
				err = results.Err()
				check.DbErr(err)
			}
		}
	}
	return resultsSlice
}
