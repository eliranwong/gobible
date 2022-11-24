package main

import (
	"fmt"
	"os"

	"github.com/eliranwong/gobible/internal/fyne"
	"github.com/eliranwong/gobible/internal/terminal"
)

func main() {
	args := os.Args
	mode := ""
	if len(args) > 1 {
		mode = args[1]
	}
	if mode == "fyne" {
		fyne.Fyne()
	} else {
		terminal.Terminal()
	}
	tidyUp()
}

func tidyUp() {
	fmt.Println("Saving configurations ...")
	// will save settings here
	fmt.Println("Closed!")
}
