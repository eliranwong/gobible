package bible

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/eliranwong/gobible/internal/check"
	"github.com/eliranwong/gobible/internal/parser"
	"github.com/eliranwong/gobible/internal/share"
)

// search bible for words in verses
// supports sql wildcards https://www.w3schools.com/sql/sql_wildcards.asp
// e.g. 'God%love' searches for verses containing 'God' followed by 'love'
func SimpleSearch(module, pattern string) {
	condition := fmt.Sprintf(`Scripture LIKE "%[1]v%[2]v%[1]v"`, "%", strings.TrimSpace(pattern))
	Search(module, condition)
}

// search bible with regular expression
// e.g. 'God.*?love' searches for verses containing 'God' followed by 'love'
func RegexpSearch(module, pattern string) {
	Search(module, fmt.Sprintf(`regexpSelect(SCRIPTURE, "%v", %v)`, pattern, share.SearchCaseSensitive))
}

// search bible with AND combination pattern
// e.g. 'God|love' searches for verses containing both 'God' and 'love'
func AndSearch(module, pattern string) {
	CombinedSearch(module, pattern, "AND")
}

// search bible with OR combination pattern
// e.g. 'God|love' searches for verses containing either 'God' or 'love'
func OrSearch(module, pattern string) {
	CombinedSearch(module, pattern, "OR")
}

func CombinedSearch(module, pattern, keyword string) {
	conditions := ""
	for i, v := range strings.Split(pattern, "|") {
		if !(i == 0) {
			conditions += fmt.Sprintf(" %v ", keyword)
		}
		conditions += fmt.Sprintf(`Scripture LIKE "%[1]v%[2]v%[1]v"`, "%", strings.TrimSpace(v))
	}
	Search(module, conditions)
}

// search bible with custom WHERE statement
// e.g. SCRIPTURE LIKE "%God%" AND SCRIPTURE LIKE "%love%"
// users can also use regular expression
// For example, to search for verses both containing words with "God" followed by "love" and containing a word "Jesus", enter:
// > .search
// > advanced
// > regexp(SCRIPTURE, "God.*?love", true) AND SCRIPTURE LIKE "%Jesus%"
// in which "regexp" is a registered function to support searching sqlite files with regular expression
// func Regexp(text string, pattern string, caseSensitive bool) bool
// text: source text
// pattern: regular expression pattern
// caseSensitive: determine if the search is case-sensitive
// Remarks: For a quicker search, select "regexp" as search method instead of "advanced" if you have only a single pattern of regular expression.
func Search(module, conditions string) {
	if share.Mode == "" {
		share.Divider()
	}
	db := getDb(module)
	defer db.Close()

	ch1 := make(chan *sql.Rows)
	ch2 := make(chan *sql.Rows)
	ch3 := make(chan *sql.Rows)
	ch4 := make(chan *sql.Rows)
	ch5 := make(chan *sql.Rows)
	ch6 := make(chan *sql.Rows)
	ch7 := make(chan *sql.Rows)
	ch8 := make(chan *sql.Rows)
	ch9 := make(chan *sql.Rows)
	ch10 := make(chan *sql.Rows)
	ch11 := make(chan *sql.Rows)
	ch12 := make(chan *sql.Rows)
	ch13 := make(chan *sql.Rows)
	ch14 := make(chan *sql.Rows)
	ch15 := make(chan *sql.Rows)
	ch16 := make(chan *sql.Rows)
	ch17 := make(chan *sql.Rows)
	ch18 := make(chan *sql.Rows)
	ch19 := make(chan *sql.Rows)
	ch20 := make(chan *sql.Rows)
	ch21 := make(chan *sql.Rows)
	ch22 := make(chan *sql.Rows)
	ch23 := make(chan *sql.Rows)
	ch24 := make(chan *sql.Rows)
	ch25 := make(chan *sql.Rows)
	ch26 := make(chan *sql.Rows)
	ch27 := make(chan *sql.Rows)
	ch28 := make(chan *sql.Rows)
	ch29 := make(chan *sql.Rows)
	ch30 := make(chan *sql.Rows)
	ch31 := make(chan *sql.Rows)
	ch32 := make(chan *sql.Rows)
	ch33 := make(chan *sql.Rows)
	ch34 := make(chan *sql.Rows)
	ch35 := make(chan *sql.Rows)
	ch36 := make(chan *sql.Rows)
	ch37 := make(chan *sql.Rows)
	ch38 := make(chan *sql.Rows)
	ch39 := make(chan *sql.Rows)
	ch40 := make(chan *sql.Rows)
	ch41 := make(chan *sql.Rows)
	ch42 := make(chan *sql.Rows)
	ch43 := make(chan *sql.Rows)
	ch44 := make(chan *sql.Rows)
	ch45 := make(chan *sql.Rows)
	ch46 := make(chan *sql.Rows)
	ch47 := make(chan *sql.Rows)
	ch48 := make(chan *sql.Rows)
	ch49 := make(chan *sql.Rows)
	ch50 := make(chan *sql.Rows)
	ch51 := make(chan *sql.Rows)
	ch52 := make(chan *sql.Rows)
	ch53 := make(chan *sql.Rows)
	ch54 := make(chan *sql.Rows)
	ch55 := make(chan *sql.Rows)
	ch56 := make(chan *sql.Rows)
	ch57 := make(chan *sql.Rows)
	ch58 := make(chan *sql.Rows)
	ch59 := make(chan *sql.Rows)
	ch60 := make(chan *sql.Rows)
	ch61 := make(chan *sql.Rows)
	ch62 := make(chan *sql.Rows)
	ch63 := make(chan *sql.Rows)
	ch64 := make(chan *sql.Rows)
	ch65 := make(chan *sql.Rows)
	ch66 := make(chan *sql.Rows)
	ch67 := make(chan *sql.Rows)
	for i := 1; i <= 66; i++ {
		go func(book int) {
			var query string
			if strings.HasPrefix(conditions, `regexpSelect(SCRIPTURE, "`) {
				query = fmt.Sprintf("SELECT Book, Chapter, Verse, %v FROM Verses WHERE Book=%v ORDER BY Book, Chapter, Verse", conditions, book)
			} else {
				query = fmt.Sprintf("SELECT * FROM Verses WHERE %v AND Book=%v ORDER BY Book, Chapter, Verse", conditions, book)
			}
			results, err := db.Query(query)
			check.DbErr(err)
			switch book {
			case 1:
				ch1 <- results
			case 2:
				ch2 <- results
			case 3:
				ch3 <- results
			case 4:
				ch4 <- results
			case 5:
				ch5 <- results
			case 6:
				ch6 <- results
			case 7:
				ch7 <- results
			case 8:
				ch8 <- results
			case 9:
				ch9 <- results
			case 10:
				ch10 <- results
			case 11:
				ch11 <- results
			case 12:
				ch12 <- results
			case 13:
				ch13 <- results
			case 14:
				ch14 <- results
			case 15:
				ch15 <- results
			case 16:
				ch16 <- results
			case 17:
				ch17 <- results
			case 18:
				ch18 <- results
			case 19:
				ch19 <- results
			case 20:
				ch20 <- results
			case 21:
				ch21 <- results
			case 22:
				ch22 <- results
			case 23:
				ch23 <- results
			case 24:
				ch24 <- results
			case 25:
				ch25 <- results
			case 26:
				ch26 <- results
			case 27:
				ch27 <- results
			case 28:
				ch28 <- results
			case 29:
				ch29 <- results
			case 30:
				ch30 <- results
			case 31:
				ch31 <- results
			case 32:
				ch32 <- results
			case 33:
				ch33 <- results
			case 34:
				ch34 <- results
			case 35:
				ch35 <- results
			case 36:
				ch36 <- results
			case 37:
				ch37 <- results
			case 38:
				ch38 <- results
			case 39:
				ch39 <- results
			case 40:
				ch40 <- results
			case 41:
				ch41 <- results
			case 42:
				ch42 <- results
			case 43:
				ch43 <- results
			case 44:
				ch44 <- results
			case 45:
				ch45 <- results
			case 46:
				ch46 <- results
			case 47:
				ch47 <- results
			case 48:
				ch48 <- results
			case 49:
				ch49 <- results
			case 50:
				ch50 <- results
			case 51:
				ch51 <- results
			case 52:
				ch52 <- results
			case 53:
				ch53 <- results
			case 54:
				ch54 <- results
			case 55:
				ch55 <- results
			case 56:
				ch56 <- results
			case 57:
				ch57 <- results
			case 58:
				ch58 <- results
			case 59:
				ch59 <- results
			case 60:
				ch60 <- results
			case 61:
				ch61 <- results
			case 62:
				ch62 <- results
			case 63:
				ch63 <- results
			case 64:
				ch64 <- results
			case 65:
				ch65 <- results
			case 66:
				ch66 <- results
			}
		}(i)
	}
	go func() {
		query := fmt.Sprintf("SELECT DISTINCT * FROM Verses WHERE %v AND Book>66 ORDER BY Book, Chapter, Verse", conditions)
		results, err := db.Query(query)
		check.DbErr(err)
		ch67 <- results
	}()
	for i := 1; i <= 67; i++ {
		select {
		case results1 := <-ch1:
			share.Ch1 <- processSearchResults(module, results1)
		case results2 := <-ch2:
			share.Ch2 <- processSearchResults(module, results2)
		case results3 := <-ch3:
			share.Ch3 <- processSearchResults(module, results3)
		case results4 := <-ch4:
			share.Ch4 <- processSearchResults(module, results4)
		case results5 := <-ch5:
			share.Ch5 <- processSearchResults(module, results5)
		case results6 := <-ch6:
			share.Ch6 <- processSearchResults(module, results6)
		case results7 := <-ch7:
			share.Ch7 <- processSearchResults(module, results7)
		case results8 := <-ch8:
			share.Ch8 <- processSearchResults(module, results8)
		case results9 := <-ch9:
			share.Ch9 <- processSearchResults(module, results9)
		case results10 := <-ch10:
			share.Ch10 <- processSearchResults(module, results10)
		case results11 := <-ch11:
			share.Ch11 <- processSearchResults(module, results11)
		case results12 := <-ch12:
			share.Ch12 <- processSearchResults(module, results12)
		case results13 := <-ch13:
			share.Ch13 <- processSearchResults(module, results13)
		case results14 := <-ch14:
			share.Ch14 <- processSearchResults(module, results14)
		case results15 := <-ch15:
			share.Ch15 <- processSearchResults(module, results15)
		case results16 := <-ch16:
			share.Ch16 <- processSearchResults(module, results16)
		case results17 := <-ch17:
			share.Ch17 <- processSearchResults(module, results17)
		case results18 := <-ch18:
			share.Ch18 <- processSearchResults(module, results18)
		case results19 := <-ch19:
			share.Ch19 <- processSearchResults(module, results19)
		case results20 := <-ch20:
			share.Ch20 <- processSearchResults(module, results20)
		case results21 := <-ch21:
			share.Ch21 <- processSearchResults(module, results21)
		case results22 := <-ch22:
			share.Ch22 <- processSearchResults(module, results22)
		case results23 := <-ch23:
			share.Ch23 <- processSearchResults(module, results23)
		case results24 := <-ch24:
			share.Ch24 <- processSearchResults(module, results24)
		case results25 := <-ch25:
			share.Ch25 <- processSearchResults(module, results25)
		case results26 := <-ch26:
			share.Ch26 <- processSearchResults(module, results26)
		case results27 := <-ch27:
			share.Ch27 <- processSearchResults(module, results27)
		case results28 := <-ch28:
			share.Ch28 <- processSearchResults(module, results28)
		case results29 := <-ch29:
			share.Ch29 <- processSearchResults(module, results29)
		case results30 := <-ch30:
			share.Ch30 <- processSearchResults(module, results30)
		case results31 := <-ch31:
			share.Ch31 <- processSearchResults(module, results31)
		case results32 := <-ch32:
			share.Ch32 <- processSearchResults(module, results32)
		case results33 := <-ch33:
			share.Ch33 <- processSearchResults(module, results33)
		case results34 := <-ch34:
			share.Ch34 <- processSearchResults(module, results34)
		case results35 := <-ch35:
			share.Ch35 <- processSearchResults(module, results35)
		case results36 := <-ch36:
			share.Ch36 <- processSearchResults(module, results36)
		case results37 := <-ch37:
			share.Ch37 <- processSearchResults(module, results37)
		case results38 := <-ch38:
			share.Ch38 <- processSearchResults(module, results38)
		case results39 := <-ch39:
			share.Ch39 <- processSearchResults(module, results39)
		case results40 := <-ch40:
			share.Ch40 <- processSearchResults(module, results40)
		case results41 := <-ch41:
			share.Ch41 <- processSearchResults(module, results41)
		case results42 := <-ch42:
			share.Ch42 <- processSearchResults(module, results42)
		case results43 := <-ch43:
			share.Ch43 <- processSearchResults(module, results43)
		case results44 := <-ch44:
			share.Ch44 <- processSearchResults(module, results44)
		case results45 := <-ch45:
			share.Ch45 <- processSearchResults(module, results45)
		case results46 := <-ch46:
			share.Ch46 <- processSearchResults(module, results46)
		case results47 := <-ch47:
			share.Ch47 <- processSearchResults(module, results47)
		case results48 := <-ch48:
			share.Ch48 <- processSearchResults(module, results48)
		case results49 := <-ch49:
			share.Ch49 <- processSearchResults(module, results49)
		case results50 := <-ch50:
			share.Ch50 <- processSearchResults(module, results50)
		case results51 := <-ch51:
			share.Ch51 <- processSearchResults(module, results51)
		case results52 := <-ch52:
			share.Ch52 <- processSearchResults(module, results52)
		case results53 := <-ch53:
			share.Ch53 <- processSearchResults(module, results53)
		case results54 := <-ch54:
			share.Ch54 <- processSearchResults(module, results54)
		case results55 := <-ch55:
			share.Ch55 <- processSearchResults(module, results55)
		case results56 := <-ch56:
			share.Ch56 <- processSearchResults(module, results56)
		case results57 := <-ch57:
			share.Ch57 <- processSearchResults(module, results57)
		case results58 := <-ch58:
			share.Ch58 <- processSearchResults(module, results58)
		case results59 := <-ch59:
			share.Ch59 <- processSearchResults(module, results59)
		case results60 := <-ch60:
			share.Ch60 <- processSearchResults(module, results60)
		case results61 := <-ch61:
			share.Ch61 <- processSearchResults(module, results61)
		case results62 := <-ch62:
			share.Ch62 <- processSearchResults(module, results62)
		case results63 := <-ch63:
			share.Ch63 <- processSearchResults(module, results63)
		case results64 := <-ch64:
			share.Ch64 <- processSearchResults(module, results64)
		case results65 := <-ch65:
			share.Ch65 <- processSearchResults(module, results65)
		case results66 := <-ch66:
			share.Ch66 <- processSearchResults(module, results66)
		case results67 := <-ch67:
			share.Ch67 <- processSearchResults(module, results67)
		}
	}
}

func processSearchResults(module string, results *sql.Rows) [][]string {
	defer results.Close()

	var resultSlice [][]string
	var err error

	for results.Next() {
		var b, c, v int
		var text string
		err = results.Scan(&b, &c, &v, &text)
		check.DbErr(err)
		text = formatVerseText(strings.TrimSpace(text))
		if !(text == "") {
			resultSlice = append(resultSlice, []string{module, parser.BcvToVerseReference([]int{b, c, v}), text})
		}
	}
	err = results.Err()
	check.DbErr(err)

	return resultSlice
}
