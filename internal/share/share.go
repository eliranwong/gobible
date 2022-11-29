// share elements
package share

import (
	"fmt"

	"fyne.io/fyne/v2"
)

var Gobible fyne.App
var Bible string
var BookName string
var BookAbb string
var Reference string
var Book int = 43
var Chapter int = 3
var Verse int = 16

var Mode string = ""
var DividerStr string = "--------------------"

func Init(gobible fyne.App) {
	Gobible = gobible
	Bible = Gobible.Preferences().StringWithFallback("bible", "NET")
	BookName = Gobible.Preferences().StringWithFallback("book_name", "John")
	BookAbb = Gobible.Preferences().StringWithFallback("book_abb", "John")
	Reference = Gobible.Preferences().StringWithFallback("reference", "John 3:16")
	Book = Gobible.Preferences().IntWithFallback("reference", 43)
	Chapter = Gobible.Preferences().IntWithFallback("chapter", 3)
	Verse = Gobible.Preferences().IntWithFallback("verse", 16)
}

func SetBible(value string) {
	Bible = value
	Gobible.Preferences().SetString("bible", value)
}

func SetBookName(value string) {
	BookName = value
	Gobible.Preferences().SetString("book_name", value)
}

func Divider() {
	fmt.Println(DividerStr)
}
