package main

import (
	"fmt"
	"strings"

	"github.com/eliranwong/gobible/internal/bible"
	//"github.com/eliranwong/gobible/internal/parser"
)

func main() {
	fmt.Println("... gobible is in development ...")
	//parser.Init()
	var bibleModule string
	var b, c, v int
	fmt.Print("Enter Bible Module:\n(e.g. KJV)\n>>> ")
	fmt.Scanln(&bibleModule)
	fmt.Print("Enter bible reference:\n(e.g. 43.3 or 43.3.16)\n>>> ")
	fmt.Scanf("%v.%v.%v", &b, &c, &v)
	bible.SingleVerse(strings.TrimSpace(bibleModule), b, c, v)
}
