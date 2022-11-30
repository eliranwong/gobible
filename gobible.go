package main

import (
	"fmt"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/eliranwong/gobible/internal/gui"
	"github.com/eliranwong/gobible/internal/share"
	"github.com/eliranwong/gobible/internal/terminal"
)

var Gobible fyne.App

func Init() {
	// set default font
	os.Setenv("FYNE_FONT", filepath.FromSlash("fonts/fonts.ttf"))
	// set appication size with FYNE_SCALE
	// read https://developer.fyne.io/architecture/scaling
	os.Setenv("FYNE_SCALE", "1.2")

	Gobible = app.NewWithID("app.uniquebible.go")

	fyneTheme := Gobible.Preferences().StringWithFallback("fyne_theme", "dark")
	if fyneTheme == "dark" {
		Gobible.Settings().SetTheme(theme.DarkTheme())
	} else {
		Gobible.Settings().SetTheme(theme.LightTheme())
	}
}

func main() {

	Init()

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
