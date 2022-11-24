package fyne

import (
	"os"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func Fyne() {
	// set default theme: "dark" or "light"
	// TODO: save in config settings later
	os.Setenv("FYNE_THEME", "dark")
	gobible := app.New()
	window := gobible.NewWindow("Go Bible")

	window.SetContent(widget.NewLabel("Go Bible!"))
	window.ShowAndRun()
}
