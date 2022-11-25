package fyne

import (
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/eliranwong/gobible/internal/bible"
	"github.com/eliranwong/gobible/internal/parser"
	"github.com/eliranwong/gobible/internal/share"
	"github.com/eliranwong/gobible/internal/shortcuts"
)

var Window fyne.Window

func Fyne() {
	// set default theme: "dark" or "light"
	// TODO: save in config settings later
	os.Setenv("FYNE_THEME", "dark")
	os.Setenv("FYNE_FONT", filepath.FromSlash("fonts/sileot.ttf"))
	gobible := app.New()
	Window = gobible.NewWindow("Go Bible")
	Window.Resize(fyne.NewSize(1024, 768))

	textArea := widget.NewEntry()
	textArea.MultiLine = true
	textArea.Wrapping = fyne.TextWrapWord

	command := widget.NewEntry()
	command.SetPlaceHolder("Enter bible reference or search item here ...")
	command.OnSubmitted = func(s string) {
		RunCommand(command.Text, share.Bible, textArea)
	}

	bibles, _ := shortcuts.WalkMatch(filepath.FromSlash("data/bibles"), "*.bible", true)
	bibleList := widget.NewList(
		func() int {
			return len(bibles)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(bibles[i])
		})

	bibleList.OnSelected = func(id widget.ListItemID) {
		share.Bible = bibles[id]
		RunCommand(command.Text, share.Bible, textArea)
	}
	content := container.NewBorder(command, nil, bibleList, nil, textArea)
	Window.SetContent(content)

	startupCommand := "John 3:16-16"
	command.Text = startupCommand
	RunCommand(startupCommand, share.Bible, textArea)
	Window.ShowAndRun()
}

func RunCommand(command, bibleModule string, textArea *widget.Entry) {
	if !(strings.TrimSpace(command) == "") {
		// reset bible text for display
		bible.Display = ""
		// parse bible reference
		references := parser.ExtractAllReferences(command, false)
		// search bible when there is no valid bible reference
		if len(references) == 0 {
			bible.AndSearch(bibleModule, command)
		} else {
			bible.Read(bibleModule, references)
		}
		// display bible text
		textArea.SetText(bible.Display)
	}
}
