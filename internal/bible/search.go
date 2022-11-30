package bible

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/eliranwong/gobible/internal/check"
	"github.com/eliranwong/gobible/internal/parser"
	"github.com/eliranwong/gobible/internal/share"
)

// search bible with AND combination pattern, e.g. Jesus|love
func AndSearch(module, pattern string) {
	//" AND ".join(['Scripture LIKE "%{0}%"'.format(m.strip()) for m in commandList[index].split("|")])
	conditions := ""
	for i, v := range strings.Split(pattern, "|") {
		if !(i == 0) {
			conditions += " AND "
		}
		conditions += fmt.Sprintf(`Scripture LIKE "%[1]v%[2]v%[1]v"`, "%", strings.TrimSpace(v))
	}
	Search(module, conditions)
}

// search bible
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
	//query := fmt.Sprintf("SELECT DISTINCT * FROM Verses WHERE %v ORDER BY Book, Chapter, Verse", conditions)
	for i := 1; i <= 66; i++ {
		go func(book int) {
			query := fmt.Sprintf("SELECT DISTINCT * FROM Verses WHERE %v AND Book=%v ORDER BY Book, Chapter, Verse", conditions, book)
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
	if share.Mode == "fyne" {
		// populate tabs as soon as content is available
		for i := 1; i <= 67; i++ {
			select {
			case results1 := <-ch1:
				share.Ch1 <- processSearchResults(results1)
			case results2 := <-ch2:
				share.Ch2 <- processSearchResults(results2)
			case results3 := <-ch3:
				share.Ch3 <- processSearchResults(results3)
			case results4 := <-ch4:
				share.Ch4 <- processSearchResults(results4)
			case results5 := <-ch5:
				share.Ch5 <- processSearchResults(results5)
			case results6 := <-ch6:
				share.Ch6 <- processSearchResults(results6)
			case results7 := <-ch7:
				share.Ch7 <- processSearchResults(results7)
			case results8 := <-ch8:
				share.Ch8 <- processSearchResults(results8)
			case results9 := <-ch9:
				share.Ch9 <- processSearchResults(results9)
			case results10 := <-ch10:
				share.Ch10 <- processSearchResults(results10)
			case results11 := <-ch11:
				share.Ch11 <- processSearchResults(results11)
			case results12 := <-ch12:
				share.Ch12 <- processSearchResults(results12)
			case results13 := <-ch13:
				share.Ch13 <- processSearchResults(results13)
			case results14 := <-ch14:
				share.Ch14 <- processSearchResults(results14)
			case results15 := <-ch15:
				share.Ch15 <- processSearchResults(results15)
			case results16 := <-ch16:
				share.Ch16 <- processSearchResults(results16)
			case results17 := <-ch17:
				share.Ch17 <- processSearchResults(results17)
			case results18 := <-ch18:
				share.Ch18 <- processSearchResults(results18)
			case results19 := <-ch19:
				share.Ch19 <- processSearchResults(results19)
			case results20 := <-ch20:
				share.Ch20 <- processSearchResults(results20)
			case results21 := <-ch21:
				share.Ch21 <- processSearchResults(results21)
			case results22 := <-ch22:
				share.Ch22 <- processSearchResults(results22)
			case results23 := <-ch23:
				share.Ch23 <- processSearchResults(results23)
			case results24 := <-ch24:
				share.Ch24 <- processSearchResults(results24)
			case results25 := <-ch25:
				share.Ch25 <- processSearchResults(results25)
			case results26 := <-ch26:
				share.Ch26 <- processSearchResults(results26)
			case results27 := <-ch27:
				share.Ch27 <- processSearchResults(results27)
			case results28 := <-ch28:
				share.Ch28 <- processSearchResults(results28)
			case results29 := <-ch29:
				share.Ch29 <- processSearchResults(results29)
			case results30 := <-ch30:
				share.Ch30 <- processSearchResults(results30)
			case results31 := <-ch31:
				share.Ch31 <- processSearchResults(results31)
			case results32 := <-ch32:
				share.Ch32 <- processSearchResults(results32)
			case results33 := <-ch33:
				share.Ch33 <- processSearchResults(results33)
			case results34 := <-ch34:
				share.Ch34 <- processSearchResults(results34)
			case results35 := <-ch35:
				share.Ch35 <- processSearchResults(results35)
			case results36 := <-ch36:
				share.Ch36 <- processSearchResults(results36)
			case results37 := <-ch37:
				share.Ch37 <- processSearchResults(results37)
			case results38 := <-ch38:
				share.Ch38 <- processSearchResults(results38)
			case results39 := <-ch39:
				share.Ch39 <- processSearchResults(results39)
			case results40 := <-ch40:
				share.Ch40 <- processSearchResults(results40)
			case results41 := <-ch41:
				share.Ch41 <- processSearchResults(results41)
			case results42 := <-ch42:
				share.Ch42 <- processSearchResults(results42)
			case results43 := <-ch43:
				share.Ch43 <- processSearchResults(results43)
			case results44 := <-ch44:
				share.Ch44 <- processSearchResults(results44)
			case results45 := <-ch45:
				share.Ch45 <- processSearchResults(results45)
			case results46 := <-ch46:
				share.Ch46 <- processSearchResults(results46)
			case results47 := <-ch47:
				share.Ch47 <- processSearchResults(results47)
			case results48 := <-ch48:
				share.Ch48 <- processSearchResults(results48)
			case results49 := <-ch49:
				share.Ch49 <- processSearchResults(results49)
			case results50 := <-ch50:
				share.Ch50 <- processSearchResults(results50)
			case results51 := <-ch51:
				share.Ch51 <- processSearchResults(results51)
			case results52 := <-ch52:
				share.Ch52 <- processSearchResults(results52)
			case results53 := <-ch53:
				share.Ch53 <- processSearchResults(results53)
			case results54 := <-ch54:
				share.Ch54 <- processSearchResults(results54)
			case results55 := <-ch55:
				share.Ch55 <- processSearchResults(results55)
			case results56 := <-ch56:
				share.Ch56 <- processSearchResults(results56)
			case results57 := <-ch57:
				share.Ch57 <- processSearchResults(results57)
			case results58 := <-ch58:
				share.Ch58 <- processSearchResults(results58)
			case results59 := <-ch59:
				share.Ch59 <- processSearchResults(results59)
			case results60 := <-ch60:
				share.Ch60 <- processSearchResults(results60)
			case results61 := <-ch61:
				share.Ch61 <- processSearchResults(results61)
			case results62 := <-ch62:
				share.Ch62 <- processSearchResults(results62)
			case results63 := <-ch63:
				share.Ch63 <- processSearchResults(results63)
			case results64 := <-ch64:
				share.Ch64 <- processSearchResults(results64)
			case results65 := <-ch65:
				share.Ch65 <- processSearchResults(results65)
			case results66 := <-ch66:
				share.Ch66 <- processSearchResults(results66)
			case results67 := <-ch67:
				share.Ch67 <- processSearchResults(results67)
			}
		}
	} else {
		// put in order
		share.Divider()
		processResults(<-ch1)
		share.Divider()
		processResults(<-ch2)
		share.Divider()
		processResults(<-ch3)
		share.Divider()
		processResults(<-ch4)
		share.Divider()
		processResults(<-ch5)
		share.Divider()
		processResults(<-ch6)
		share.Divider()
		processResults(<-ch7)
		share.Divider()
		processResults(<-ch8)
		share.Divider()
		processResults(<-ch9)
		share.Divider()
		processResults(<-ch10)
		share.Divider()
		processResults(<-ch11)
		share.Divider()
		processResults(<-ch12)
		share.Divider()
		processResults(<-ch13)
		share.Divider()
		processResults(<-ch14)
		share.Divider()
		processResults(<-ch15)
		share.Divider()
		processResults(<-ch16)
		share.Divider()
		processResults(<-ch17)
		share.Divider()
		processResults(<-ch18)
		share.Divider()
		processResults(<-ch19)
		share.Divider()
		processResults(<-ch20)
		share.Divider()
		processResults(<-ch21)
		share.Divider()
		processResults(<-ch22)
		share.Divider()
		processResults(<-ch23)
		share.Divider()
		processResults(<-ch24)
		share.Divider()
		processResults(<-ch25)
		share.Divider()
		processResults(<-ch26)
		share.Divider()
		processResults(<-ch27)
		share.Divider()
		processResults(<-ch28)
		share.Divider()
		processResults(<-ch29)
		share.Divider()
		processResults(<-ch30)
		share.Divider()
		processResults(<-ch31)
		share.Divider()
		processResults(<-ch32)
		share.Divider()
		processResults(<-ch33)
		share.Divider()
		processResults(<-ch34)
		share.Divider()
		processResults(<-ch35)
		share.Divider()
		processResults(<-ch36)
		share.Divider()
		processResults(<-ch37)
		share.Divider()
		processResults(<-ch38)
		share.Divider()
		processResults(<-ch39)
		share.Divider()
		processResults(<-ch40)
		share.Divider()
		processResults(<-ch41)
		share.Divider()
		processResults(<-ch42)
		share.Divider()
		processResults(<-ch43)
		share.Divider()
		processResults(<-ch44)
		share.Divider()
		processResults(<-ch45)
		share.Divider()
		processResults(<-ch46)
		share.Divider()
		processResults(<-ch47)
		share.Divider()
		processResults(<-ch48)
		share.Divider()
		processResults(<-ch49)
		share.Divider()
		processResults(<-ch50)
		share.Divider()
		processResults(<-ch51)
		share.Divider()
		processResults(<-ch52)
		share.Divider()
		processResults(<-ch53)
		share.Divider()
		processResults(<-ch54)
		share.Divider()
		processResults(<-ch55)
		share.Divider()
		processResults(<-ch56)
		share.Divider()
		processResults(<-ch57)
		share.Divider()
		processResults(<-ch58)
		share.Divider()
		processResults(<-ch59)
		share.Divider()
		processResults(<-ch60)
		share.Divider()
		processResults(<-ch61)
		share.Divider()
		processResults(<-ch62)
		share.Divider()
		processResults(<-ch63)
		share.Divider()
		processResults(<-ch64)
		share.Divider()
		processResults(<-ch65)
		share.Divider()
		processResults(<-ch66)
		share.Divider()
		processResults(<-ch67)
	}
}

func processSearchResults(results *sql.Rows) string {
	defer results.Close()

	var searchDisplay string = ""
	var err error
	total := 0

	for results.Next() {
		var b, c, v int
		var text string
		err = results.Scan(&b, &c, &v, &text)
		check.DbErr(err)
		text = formatVerseText(text)
		display := fmt.Sprintf("%v %v", parser.BcvToVerseReference([]int{b, c, v}), text)
		searchDisplay += display
		searchDisplay += "\n"
		total += 1
	}
	err = results.Err()
	check.DbErr(err)

	return fmt.Sprintf("[total of %v verse(s)]\n%v", total, searchDisplay)
}
