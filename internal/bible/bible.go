// Package bible
// retrieve bible module data
package bible

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/eliranwong/gobible/internal/check"
	"github.com/eliranwong/gobible/internal/parser"
	"github.com/eliranwong/gobible/internal/share"
	//_ "github.com/mattn/go-sqlite3"
)

var Chapter strings.Builder

func getDb(module string) *sql.DB {
	dbPath := filepath.Join(share.Data, filepath.FromSlash(fmt.Sprintf("bibles/%v.bible", module)))
	if check.FileExists(dbPath) {
		share.Bible = module
	} else {
		dbPath = filepath.Join(share.Data, filepath.FromSlash(fmt.Sprintf("bibles/%v.bible", share.Bible)))
	}
	//dbPath := filepath.FromSlash(filePath)
	//db, err := sql.Open("sqlite3", dbPath)
	var dataSourceName string
	if share.SearchCaseSensitive {
		dataSourceName = fmt.Sprintf("%v?_case_sensitive_like=1", dbPath)
	} else {
		dataSourceName = dbPath
	}
	db, err := sql.Open("sqlite3_custom", dataSourceName)
	check.DbErr(err)
	return db
}

// read bible reference
func Read(module string, references [][]int) {
	if len(references) > 0 {
		updateSelection(references[0])
		if share.Mode == "" {
			share.Divider()
		}
		if len(references) == 1 && len(references[0]) == 3 {
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
}

func updateSelection(bcv []int) {
	share.Book = bcv[0]
	share.BookName = parser.BookNumberToName(share.Book)
	share.BookAbb = parser.BookNumberToAbb(share.Book)
	share.Chapter = bcv[1]
	share.Verse = bcv[2]
	share.Reference = fmt.Sprintf(`%v %v:%v`, share.BookAbb, share.Chapter, share.Verse)
}

func ReadChapter(module string, bcv []int) {
	ReadSingle(module, bcv)
	Chapter.WriteString(share.DividerStr)
	Chapter.WriteString("\n")
	ReadSingle(module, []int{bcv[0], bcv[1], 0})
}

// read either single verse or single chapter
func ReadSingle(module string, bcv []int) {
	db := getDb(module)
	defer db.Close()

	b, c, v := bcv[0], bcv[1], bcv[2]

	if b == 0 {
		b = 1
	}
	if c == 0 {
		c = 1
	}
	query := ""
	if v == 0 {
		query = fmt.Sprintf("SELECT * from Verses WHERE Book=%v AND Chapter=%v ORDER BY Verse", b, c)
	} else {
		query = fmt.Sprintf("SELECT * from Verses WHERE Book=%v AND Chapter=%v AND Verse=%v", b, c, v)
	}
	results, err := db.Query(query)
	check.DbErr(err)
	processResults(results)
}

// read multiple verses
func ReadMultiple(module string, bcv []int) {
	var b, c, cs, vs, ce, ve int
	b, cs, vs = bcv[0], bcv[1], bcv[2]

	if len(bcv) == 4 {
		ve = bcv[3]
		displayChapterVerses(module, b, cs, vs, ve)
	} else if len(bcv) == 5 {
		ce, ve = bcv[3], bcv[4]

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
	processResults(results)
}

// process sqlite query results
func processResults(results *sql.Rows) {
	defer results.Close()

	var err error
	total := 0

	for results.Next() {
		var b, c, v int
		var text string
		err = results.Scan(&b, &c, &v, &text)
		check.DbErr(err)
		text = formatVerseText(text)
		display := fmt.Sprintf("%v %v\n", share.Info(parser.BcvToVerseReference([]int{b, c, v})), text)
		Chapter.WriteString(display)
		total += 1
	}
	err = results.Err()
	check.DbErr(err)

	// show total verses
	if share.Mode == "" && total > 1 {
		message := fmt.Sprintf("[total of %v verse(s)]", total)
		Chapter.WriteString(message)
	}
}

func formatVerseText(text string) string {
	text = strings.ReplaceAll(text, "<gloss>", " ")
	text = regexp.MustCompile("<[^<>]*?>").ReplaceAllString(text, "")
	return strings.TrimSpace(text)
}

func GetBooks(module string) []int {
	var books []int = []int{}
	db := getDb(module)
	defer db.Close()
	query := "SELECT DISTINCT Book FROM Verses ORDER BY Book"
	results, err := db.Query(query)
	check.DbErr(err)

	for results.Next() {
		var b int
		err = results.Scan(&b)
		check.DbErr(err)
		books = append(books, b)
	}
	err = results.Err()
	check.DbErr(err)
	return books
}

func GetBookNames(module string) []string {
	bookNames := []string{}
	for _, book := range GetBooks(module) {
		bookNames = append(bookNames, parser.StandardBookname[strconv.Itoa(book)])
	}
	return bookNames
}

func GetBookAbbs(module string) []string {
	bookNames := []string{}
	for _, book := range GetBooks(module) {
		bookNames = append(bookNames, parser.StandardAbbreviation[strconv.Itoa(book)])
	}
	return bookNames
}

func GetBookNamesAndAbbs(module string) []string {
	bookNamesAndAbbs := []string{}
	for _, book := range GetBooks(module) {
		bookNamesAndAbbs = append(bookNamesAndAbbs, parser.StandardBookname[strconv.Itoa(book)])
		bookNamesAndAbbs = append(bookNamesAndAbbs, parser.StandardAbbreviation[strconv.Itoa(book)])
	}
	return bookNamesAndAbbs
}

func GetFirstBook(module string) int {
	books := GetBooks(module)
	if len(books) > 0 {
		return books[0]
	}
	return 0
}

func GetLastBook(module string) int {
	books := GetBooks(module)
	if len(books) > 0 {
		return books[len(books)-1]
	}
	return 0
}

func GetPreviousBook(module string, b int) int {
	b -= 1
	if b >= GetFirstBook(module) {
		return b
	}
	return GetLastBook(module)
}

func GetNextBook(module string, b int) int {
	b += 1
	if b <= GetLastBook(module) {
		return b
	}
	return GetFirstBook(module)
}

func GetChapters(module string, b int) []int {
	var chapters []int = []int{}
	db := getDb(module)
	defer db.Close()
	query := fmt.Sprintf("SELECT DISTINCT Chapter FROM Verses WHERE Book=%v ORDER BY Chapter", b)
	results, err := db.Query(query)
	check.DbErr(err)

	for results.Next() {
		var c int
		err = results.Scan(&c)
		check.DbErr(err)
		chapters = append(chapters, c)
	}
	err = results.Err()
	check.DbErr(err)
	return chapters
}

func GetFirstChapter(module string, b int) int {
	chapters := GetChapters(module, b)
	if len(chapters) > 0 {
		return chapters[0]
	}
	return 0
}

func GetLastChapter(module string, b int) int {
	chapters := GetChapters(module, b)
	if len(chapters) > 0 {
		return chapters[len(chapters)-1]
	}
	return 0
}

func GetPreviousChapter(module string, b, c int) int {
	c -= 1
	if c >= GetFirstChapter(module, b) {
		return c
	}
	return c + 1
}

func GetNextChapter(module string, b, c int) int {
	c += 1
	if c <= GetLastChapter(module, b) {
		return c
	}
	return c - 1
}

func GetPreviousChapterRef() string {
	var ref string
	c := GetPreviousChapter(share.Bible, share.Book, share.Chapter)
	if c != share.Chapter {
		ref = parser.BcvToVerseReference([]int{share.Book, c, 1})
	} else {
		b := GetPreviousBook(share.Bible, share.Book)
		c := GetLastChapter(share.Bible, b)
		ref = parser.BcvToVerseReference([]int{b, c, 1})
	}
	return ref
}

func GetNextChapterRef() string {
	var ref string
	c := GetNextChapter(share.Bible, share.Book, share.Chapter)
	if c != share.Chapter {
		ref = parser.BcvToVerseReference([]int{share.Book, c, 1})
	} else {
		b := GetNextBook(share.Bible, share.Book)
		c := GetFirstChapter(share.Bible, b)
		ref = parser.BcvToVerseReference([]int{b, c, 1})
	}
	return ref
}
