// share elements
package share

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/eliranwong/gobible/internal/check"
	"github.com/eliranwong/gobible/internal/regex"
	"github.com/eliranwong/gobible/internal/shortcuts"
	sqlite "github.com/mattn/go-sqlite3"
)

var Data string = "gobible_data"
var FyneTheme string = "dark"
var Bibles []string
var SelectedBibles []string = []string{"KJV", "NET"}
var Bible string = "NET"
var BookName string = "John"
var BookAbb string = "John"
var Reference string = "John 3:16"
var Book int = 43
var Chapter int = 3
var Verse int = 16

var Mode string = ""
var DividerStr string = "--------------------"
var SearchMethod = "and"
var SearchCaseSensitive = false

var Ch1 chan [][]string = make(chan [][]string)
var Ch2 chan [][]string = make(chan [][]string)
var Ch3 chan [][]string = make(chan [][]string)
var Ch4 chan [][]string = make(chan [][]string)
var Ch5 chan [][]string = make(chan [][]string)
var Ch6 chan [][]string = make(chan [][]string)
var Ch7 chan [][]string = make(chan [][]string)
var Ch8 chan [][]string = make(chan [][]string)
var Ch9 chan [][]string = make(chan [][]string)
var Ch10 chan [][]string = make(chan [][]string)
var Ch11 chan [][]string = make(chan [][]string)
var Ch12 chan [][]string = make(chan [][]string)
var Ch13 chan [][]string = make(chan [][]string)
var Ch14 chan [][]string = make(chan [][]string)
var Ch15 chan [][]string = make(chan [][]string)
var Ch16 chan [][]string = make(chan [][]string)
var Ch17 chan [][]string = make(chan [][]string)
var Ch18 chan [][]string = make(chan [][]string)
var Ch19 chan [][]string = make(chan [][]string)
var Ch20 chan [][]string = make(chan [][]string)
var Ch21 chan [][]string = make(chan [][]string)
var Ch22 chan [][]string = make(chan [][]string)
var Ch23 chan [][]string = make(chan [][]string)
var Ch24 chan [][]string = make(chan [][]string)
var Ch25 chan [][]string = make(chan [][]string)
var Ch26 chan [][]string = make(chan [][]string)
var Ch27 chan [][]string = make(chan [][]string)
var Ch28 chan [][]string = make(chan [][]string)
var Ch29 chan [][]string = make(chan [][]string)
var Ch30 chan [][]string = make(chan [][]string)
var Ch31 chan [][]string = make(chan [][]string)
var Ch32 chan [][]string = make(chan [][]string)
var Ch33 chan [][]string = make(chan [][]string)
var Ch34 chan [][]string = make(chan [][]string)
var Ch35 chan [][]string = make(chan [][]string)
var Ch36 chan [][]string = make(chan [][]string)
var Ch37 chan [][]string = make(chan [][]string)
var Ch38 chan [][]string = make(chan [][]string)
var Ch39 chan [][]string = make(chan [][]string)
var Ch40 chan [][]string = make(chan [][]string)
var Ch41 chan [][]string = make(chan [][]string)
var Ch42 chan [][]string = make(chan [][]string)
var Ch43 chan [][]string = make(chan [][]string)
var Ch44 chan [][]string = make(chan [][]string)
var Ch45 chan [][]string = make(chan [][]string)
var Ch46 chan [][]string = make(chan [][]string)
var Ch47 chan [][]string = make(chan [][]string)
var Ch48 chan [][]string = make(chan [][]string)
var Ch49 chan [][]string = make(chan [][]string)
var Ch50 chan [][]string = make(chan [][]string)
var Ch51 chan [][]string = make(chan [][]string)
var Ch52 chan [][]string = make(chan [][]string)
var Ch53 chan [][]string = make(chan [][]string)
var Ch54 chan [][]string = make(chan [][]string)
var Ch55 chan [][]string = make(chan [][]string)
var Ch56 chan [][]string = make(chan [][]string)
var Ch57 chan [][]string = make(chan [][]string)
var Ch58 chan [][]string = make(chan [][]string)
var Ch59 chan [][]string = make(chan [][]string)
var Ch60 chan [][]string = make(chan [][]string)
var Ch61 chan [][]string = make(chan [][]string)
var Ch62 chan [][]string = make(chan [][]string)
var Ch63 chan [][]string = make(chan [][]string)
var Ch64 chan [][]string = make(chan [][]string)
var Ch65 chan [][]string = make(chan [][]string)
var Ch66 chan [][]string = make(chan [][]string)
var Ch67 chan [][]string = make(chan [][]string)

func Divider() {
	fmt.Println(DividerStr)
}

func Check() {
	fmt.Println(FyneTheme, Bible, BookName, BookAbb, Reference, Book, Chapter, Verse)
}

func SetupData() {
	// default data directory; relative to executable file
	if !(check.FileExists(filepath.Join(Data, "bibles", "NET.bible"))) {
		wd, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		// mac binary ends with '/GoBible.app/Contents/MacOS'
		wd = strings.Replace(wd, "/GoBible.app/Contents/MacOS", "", -1)
		Data = filepath.Join(wd, "gobible_data")
	}
	// alternate path: use gobible data installed in home directory
	alternateDataPath := filepath.Join(os.Getenv("HOME"), "gobible", "gobible_data")
	if !(check.FileExists(Data)) && (check.FileExists(alternateDataPath)) {
		Data = alternateDataPath
	}
	bibles, _ := shortcuts.WalkMatch(filepath.Join(Data, filepath.FromSlash("bibles")), "*.bible", true)
	bibles = RemoveEmptyString(bibles)
	Bibles = bibles
}

func RemoveEmptyString(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func StringSliceToIntSlice(stringSlice []string) []int {
	intSlice := []int{}
	for _, v := range stringSlice {
		num, _ := strconv.Atoi(v)
		intSlice = append(intSlice, num)
	}
	return intSlice
}

func IntSliceToStringSlice(intSlice []int) []string {
	stringSlice := []string{}
	for _, v := range intSlice {
		stringSlice = append(stringSlice, strconv.Itoa(v))
	}
	return stringSlice
}

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s", name, elapsed)
}

func RegisterSql() {
	// supports query like, FROM VERSES WHERE re("%v", SCRIPTURE)
	sql.Register("sqlite3_custom", &sqlite.SQLiteDriver{
		ConnectHook: func(conn *sqlite.SQLiteConn) error {
			if err := conn.RegisterFunc("regexpSelect", regex.RegexpSelect, true); err != nil {
				return err
			}
			if err := conn.RegisterFunc("regexp", regex.Regexp, true); err != nil {
				return err
			}
			return nil
		},
	})
}
