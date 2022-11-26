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
	if share.Mode == "fyne" {
		fyne.Fyne()
	} else {
		terminal.Terminal()
	}
	tidyUp()
}

func tidyUp() {
	//TODO save config
	fmt.Println("Saving configurations ...")
	// will save settings here
	fmt.Println("Closed")
}
