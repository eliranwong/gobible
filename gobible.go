package main

import (
	"fmt"
	"os"

	"github.com/eliranwong/gobible/internal/fyne"
	"github.com/eliranwong/gobible/internal/share"
	"github.com/eliranwong/gobible/internal/terminal"
)

func main() {

	args := os.Args
	if len(args) > 1 {
		share.Mode = args[1]
	}
	// for running "fyne" for development purpose
	share.Mode = "fyne"
	if share.Mode == "fyne" {
		//fyne.Fyne()
		fyne.Fyne()
	} else {
		terminal.Terminal()
	}
	tidyUp()
}

func tidyUp() {
	fmt.Println("Closed")
}
