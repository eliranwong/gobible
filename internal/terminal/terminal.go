// terminal mode interface
package terminal

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/eliranwong/gobible/internal/bible"
	"github.com/eliranwong/gobible/internal/parser"
	"github.com/eliranwong/gobible/internal/share"
	"github.com/eliranwong/gobible/internal/shortcuts"
)

func Terminal() {
	for {
		share.Divider()
		prompt()
	}
}

func prompt() {
	var bibleModule string
	bibles, _ := shortcuts.WalkMatch(filepath.FromSlash("data/bibles"), "*.bible", true)
	fmt.Printf("Enter Bible Module:\n| %v |\n(current: %v)\n>>> ", (strings.Join(bibles, " | ")), share.Bible)
	fmt.Scanln(&bibleModule)
	fmt.Println("Read or search bible")
	fmt.Print("To read, enter bible reference(s):\n(e.g. John 3:16; Romans 5:8)\n>>> ")
	in := bufio.NewReader(os.Stdin)
	command, _ := in.ReadString('\n')
	RunCommand(command, bibleModule)
}

func RunCommand(command, bibleModule string) {
	if !(strings.TrimSpace(command) == "") {
		references := parser.ExtractAllReferences(command, false)
		// search bible when there is no valid bible reference
		if len(references) == 0 {
			bible.AndSearch(bibleModule, command)
		} else {
			bible.Read(bibleModule, references)
		}
	}
}
