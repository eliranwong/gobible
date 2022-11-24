package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/eliranwong/gobible/internal/bible"
	"github.com/eliranwong/gobible/internal/parser"
	"github.com/eliranwong/gobible/internal/shortcuts"
)

func main() {
	fmt.Println("... gobible is in development ...")
	for {
		shortcuts.Divider()
		prompt()
	}
}

func prompt() {
	var bibleModule string
	bibles, _ := shortcuts.WalkMatch(filepath.FromSlash("data/bibles"), "*.bible", true)
	fmt.Printf("Enter Bible Module:\n| %v |\n>>> ", (strings.Join(bibles, " | ")))
	fmt.Scanln(&bibleModule)
	fmt.Print("Enter bible reference:\n(e.g. John 3:16; Romans 5:8)\n>>> ")
	in := bufio.NewReader(os.Stdin)
	command, _ := in.ReadString('\n')
	references := parser.ExtractAllReferences(command, false)
	bible.Read(bibleModule, references)
}
