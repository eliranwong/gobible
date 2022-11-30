package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/eliranwong/gobible/internal/gui"
	"github.com/eliranwong/gobible/internal/share"
	"github.com/eliranwong/gobible/internal/terminal"
)

var Gobible fyne.App

func main() {
	Gobible = app.NewWithID("app.uniquebible.go")
	share.Init(Gobible)

	args := os.Args
	if len(args) > 1 {
		share.Mode = args[1]
	}
	if share.Mode == "fyne" {
		gui.Fyne(Gobible)
	} else {
		terminal.Terminal(Gobible)
	}
	tidyUp()
}

func tidyUp() {
	fmt.Println("Closed")
}
