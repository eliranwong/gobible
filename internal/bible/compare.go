package bible

import (
	"fmt"
	"sort"

	"github.com/eliranwong/gobible/internal/check"
	"github.com/eliranwong/gobible/internal/parser"
	"github.com/eliranwong/gobible/internal/share"
)

// compare bible
func CompareChapter(reference string) [][]string {
	var resultsSlice [][]string = [][]string{}
	// parse bible reference
	references := parser.ExtractAllReferences(reference, false)
	if len(references) > 0 {
		b, c := references[0][0], references[0][1]
		query := fmt.Sprintf("SELECT * from Verses WHERE Book=%v AND Chapter=%v ORDER BY Verse", b, c)

		var verseMap map[int][][]string = map[int][][]string{}
		var verseMapKeys []int = []int{}
		var resultSlice []string

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
				val, ok := verseMap[v]
				if ok {
					verseMap[v] = append(val, resultSlice)
				} else {
					verseMapKeys = append(verseMapKeys, v)
					verseMap[v] = [][]string{resultSlice}
				}
			}
			err = results.Err()
			check.DbErr(err)
		}

		sort.Ints(verseMapKeys)
		for _, verseMapKey := range verseMapKeys {
			resultsSlice = append(resultsSlice, verseMap[verseMapKey]...)
		}
	}
	return resultsSlice
}

func CompareVerse(reference string) [][]string {
	var resultsSlice [][]string = [][]string{}
	// parse bible reference
	references := parser.ExtractAllReferences(reference, false)
	if len(references) > 0 {
		b, c, v := references[0][0], references[0][1], references[0][2]
		query := fmt.Sprintf("SELECT * from Verses WHERE Book=%v AND Chapter=%v AND Verse=%v", b, c, v)

		var resultSlice []string

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
	return resultsSlice
}
