package main

import (
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
	// force to use "fyne" mode to create gui executable
	//share.Mode = "fyne"
	if share.Mode == "fyne" {
		fyne.Fyne()
	} else {
		terminal.Terminal()
	}
}
