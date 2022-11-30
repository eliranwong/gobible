package gui

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/eliranwong/gobible/internal/bible"
	"github.com/eliranwong/gobible/internal/check"
	"github.com/eliranwong/gobible/internal/parser"
	"github.com/eliranwong/gobible/internal/share"
	"github.com/eliranwong/gobible/internal/shortcuts"
	"github.com/fyne-io/terminal"
)

var Gobible fyne.App
var Window fyne.Window
var BibleTabs *container.DocTabs
var Tab0, Tab1, Tab2, Tab3, Tab4, Tab5, Tab6, Tab7, Tab8, Tab9 *widget.Entry
var Tab10, Tab11, Tab12, Tab13, Tab14, Tab15, Tab16, Tab17, Tab18, Tab19 *widget.Entry
var Tab20, Tab21, Tab22, Tab23, Tab24, Tab25, Tab26, Tab27, Tab28, Tab29 *widget.Entry
var Tab30, Tab31, Tab32, Tab33, Tab34, Tab35, Tab36, Tab37, Tab38, Tab39 *widget.Entry
var Tab40, Tab41, Tab42, Tab43, Tab44, Tab45, Tab46, Tab47, Tab48, Tab49 *widget.Entry

func config(gobible fyne.App) {
	Gobible = gobible
	theme := Gobible.Preferences().StringWithFallback("fyne_theme", "dark")
	os.Setenv("FYNE_THEME", theme)
	// set default font
	os.Setenv("FYNE_FONT", filepath.FromSlash("fonts/fonts.ttf"))
	// set appication size with FYNE_SCALE
	// read https://developer.fyne.io/architecture/scaling
	scale := Gobible.Preferences().StringWithFallback("fyne_scale", "1.2")
	os.Setenv("FYNE_SCALE", scale)

	Window = Gobible.NewWindow("Go Bible")
	Window.Resize(fyne.NewSize(1024, 768))
}

func Fyne(gobible fyne.App) {
	config(gobible)

	// tabs for displaying bible text
	bibleTabsContainer := makeDocTabsTab()
	// text entry and drowndown menu for bible selection
	bibles, _ := shortcuts.WalkMatch(filepath.FromSlash("data/bibles"), "*.bible", true)
	bibleSelect := widget.NewSelectEntry(bibles)
	bibleSelect.PlaceHolder = share.Bible
	bibleSelect.OnChanged = func(s string) {
		filePath := fmt.Sprintf("data/bibles/%v.bible", s)
		if check.FileExists(filePath) {
			share.SetBible(s)
			bibleSelect.PlaceHolder = share.Bible
			RunCommand(share.Reference, share.Bible, BibleTabs)
		}
	}
	bibleSelect.OnSubmitted = func(s string) {
		share.SetBible(s)
		RunCommand(share.Reference, share.Bible, BibleTabs)
	}
	// bible passage selection tree
	chapters := makeTree()
	// layout putting together bible navigation elements
	bibleNavigator := container.NewBorder(bibleSelect, nil, nil, nil, chapters)
	bibleNavigator.Hide()
	// bible layout
	bibleLayout := container.NewBorder(nil, nil, bibleNavigator, nil, bibleTabsContainer)

	// bible selection list
	/*
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
			share.SetBible(bibles[id])
			RunCommand(command.Text, share.Bible, BibleTabs)
		}*/

	// button to show / hide bible navigator menu
	showHideBibleNavigator := widget.NewButtonWithIcon("", theme.MenuIcon(), func() {
		if bibleNavigator.Visible() {
			bibleNavigator.Hide()
			bibleLayout.Refresh()
		} else {
			bibleNavigator.Show()
			bibleLayout.Refresh()
		}
	})
	// text entry for command
	command := widget.NewEntry()
	command.SetPlaceHolder("Enter bible reference or search item here ...")
	command.OnSubmitted = func(s string) {
		RunCommand(s, share.Bible, BibleTabs)
	}
	// top right buttons
	/* sub-menu example
	shareItem := fyne.NewMenuItem("Share via", nil)
	shareItem.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Twitter", func() { fmt.Println("context menu Share->Twitter") }),
		fyne.NewMenuItem("Reddit", func() { fmt.Println("context menu Share->Reddit") }),
	)*/
	featureMenuButton := newContextMenuButton("", theme.ContentAddIcon(), fyne.NewMenu("",
		fyne.NewMenuItem("Terminal", newTerminalWindow),
	))
	settingButton := newContextMenuButton("", theme.SettingsIcon(), fyne.NewMenu("",
		fyne.NewMenuItem("Quit", fyne.CurrentApp().Quit),
	))

	// created a text area temporarily, for filling up the right side of the split container
	textArea := widget.NewMultiLineEntry()
	textArea.Wrapping = fyne.TextWrapWord

	topRightButton := container.NewHBox(featureMenuButton, settingButton)
	mainTop := container.NewBorder(nil, nil, showHideBibleNavigator, topRightButton, command)
	mainCentre := container.NewHSplit(bibleLayout, textArea)
	mainLayout := container.NewBorder(mainTop, nil, nil, nil, mainCentre)
	Window.SetContent(mainLayout)

	startupCommand := fmt.Sprintf(`%v %v:%v`, share.BookAbb, share.Chapter, share.Verse)
	command.Text = startupCommand
	RunCommand(startupCommand, share.Bible, BibleTabs)
	Window.ShowAndRun()
}

type contextMenuButton struct {
	widget.Button
	menu *fyne.Menu
}

func (b *contextMenuButton) Tapped(e *fyne.PointEvent) {
	widget.ShowPopUpMenuAtPosition(b.menu, fyne.CurrentApp().Driver().CanvasForObject(b), e.AbsolutePosition)
}

func newContextMenuButton(label string, icon fyne.Resource, menu *fyne.Menu) *contextMenuButton {
	b := &contextMenuButton{menu: menu}
	b.Text = label
	b.Icon = icon

	b.ExtendBaseWidget(b)
	return b
}

func makeTree() fyne.CanvasObject {
	reference := ""

	data := map[string][]string{
		"": {},
	}
	var name, abb, bStr, cStr, chapterStr string
	var cTotal, vTotal int
	for b := 1; b < 67; b++ {
		bStr = strconv.Itoa(b)
		name = parser.StandardBookname[bStr]
		abb = parser.StandardAbbreviation[bStr]
		cTotal = parser.Chapters[b]
		data[""] = append(data[""], name)
		for c := 1; c <= cTotal; c++ {
			cStr = strconv.Itoa(c)
			chapterStr = fmt.Sprintf(`%v %v`, abb, cStr)
			data[name] = append(data[name], chapterStr)
			vTotal = parser.Verses[b][c]
			for v := 1; v <= vTotal; v++ {
				data[chapterStr] = append(data[chapterStr], fmt.Sprintf(`%v %v:%v`, abb, cStr, v))
			}
		}
	}

	tree := widget.NewTreeWithStrings(data)
	tree.OnBranchOpened = func(id string) {
		reference = fmt.Sprintf(`%v %v`, id, 1)
	}
	tree.OnSelected = func(id string) {
		_, err := strconv.Atoi(id[len(id)-1:])
		if err == nil {
			reference = id
		} else {
			reference = fmt.Sprintf(`%v %v`, id, 1)
		}
		RunCommand(reference, share.Bible, BibleTabs)
	}
	/*tree.OnUnselected = func(id string) {
		//
	}*/
	//tree.OpenBranch("John")
	return tree
}

func newTerminalWindow() {
	w := fyne.CurrentApp().NewWindow("Terminal")
	w.Resize(fyne.NewSize(800, 600))
	w.SetContent(makeTerminal())
	w.Show()
}

func makeTerminal() *terminal.Terminal {
	t := terminal.New()
	go func() {
		_ = t.RunLocalShell()
	}()
	return t
}

func makeDocTabsTab() fyne.CanvasObject {
	Tab0 = widget.NewMultiLineEntry()
	Tab0.Wrapping = fyne.TextWrapWord
	BibleTabs = container.NewDocTabs(
		container.NewTabItem(share.Bible, Tab0),
	)
	BibleTabs.CreateTab = func() *container.TabItem {
		i := len(BibleTabs.Items)
		switch i {
		case 0:
			Tab0 = widget.NewMultiLineEntry()
			Tab0.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab0)
		case 1:
			Tab1 = widget.NewMultiLineEntry()
			Tab1.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab1)
		case 2:
			Tab2 = widget.NewMultiLineEntry()
			Tab2.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab2)
		case 3:
			Tab3 = widget.NewMultiLineEntry()
			Tab3.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab3)
		case 4:
			Tab4 = widget.NewMultiLineEntry()
			Tab4.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab4)
		case 5:
			Tab5 = widget.NewMultiLineEntry()
			Tab5.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab5)
		case 6:
			Tab6 = widget.NewMultiLineEntry()
			Tab6.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab6)
		case 7:
			Tab7 = widget.NewMultiLineEntry()
			Tab7.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab7)
		case 8:
			Tab8 = widget.NewMultiLineEntry()
			Tab8.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab8)
		case 9:
			Tab9 = widget.NewMultiLineEntry()
			Tab9.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab9)
		case 10:
			Tab10 = widget.NewMultiLineEntry()
			Tab10.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab10)
		case 11:
			Tab11 = widget.NewMultiLineEntry()
			Tab11.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab11)
		case 12:
			Tab12 = widget.NewMultiLineEntry()
			Tab12.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab12)
		case 13:
			Tab13 = widget.NewMultiLineEntry()
			Tab13.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab13)
		case 14:
			Tab14 = widget.NewMultiLineEntry()
			Tab14.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab14)
		case 15:
			Tab15 = widget.NewMultiLineEntry()
			Tab15.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab15)
		case 16:
			Tab16 = widget.NewMultiLineEntry()
			Tab16.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab16)
		case 17:
			Tab17 = widget.NewMultiLineEntry()
			Tab17.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab17)
		case 18:
			Tab18 = widget.NewMultiLineEntry()
			Tab18.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab18)
		case 19:
			Tab19 = widget.NewMultiLineEntry()
			Tab19.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab19)
		case 20:
			Tab20 = widget.NewMultiLineEntry()
			Tab20.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab20)
		case 21:
			Tab21 = widget.NewMultiLineEntry()
			Tab21.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab21)
		case 22:
			Tab22 = widget.NewMultiLineEntry()
			Tab22.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab22)
		case 23:
			Tab23 = widget.NewMultiLineEntry()
			Tab23.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab23)
		case 24:
			Tab24 = widget.NewMultiLineEntry()
			Tab24.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab24)
		case 25:
			Tab25 = widget.NewMultiLineEntry()
			Tab25.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab25)
		case 26:
			Tab26 = widget.NewMultiLineEntry()
			Tab26.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab26)
		case 27:
			Tab27 = widget.NewMultiLineEntry()
			Tab27.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab27)
		case 28:
			Tab28 = widget.NewMultiLineEntry()
			Tab28.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab28)
		case 29:
			Tab29 = widget.NewMultiLineEntry()
			Tab29.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab29)
		case 30:
			Tab30 = widget.NewMultiLineEntry()
			Tab30.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab30)
		case 31:
			Tab31 = widget.NewMultiLineEntry()
			Tab31.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab31)
		case 32:
			Tab32 = widget.NewMultiLineEntry()
			Tab32.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab32)
		case 33:
			Tab33 = widget.NewMultiLineEntry()
			Tab33.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab33)
		case 34:
			Tab34 = widget.NewMultiLineEntry()
			Tab34.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab34)
		case 35:
			Tab35 = widget.NewMultiLineEntry()
			Tab35.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab35)
		case 36:
			Tab36 = widget.NewMultiLineEntry()
			Tab36.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab36)
		case 37:
			Tab37 = widget.NewMultiLineEntry()
			Tab37.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab37)
		case 38:
			Tab38 = widget.NewMultiLineEntry()
			Tab38.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab38)
		case 39:
			Tab39 = widget.NewMultiLineEntry()
			Tab39.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab39)
		case 40:
			Tab40 = widget.NewMultiLineEntry()
			Tab40.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab40)
		case 41:
			Tab41 = widget.NewMultiLineEntry()
			Tab41.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab41)
		case 42:
			Tab42 = widget.NewMultiLineEntry()
			Tab42.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab42)
		case 43:
			Tab43 = widget.NewMultiLineEntry()
			Tab43.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab43)
		case 44:
			Tab44 = widget.NewMultiLineEntry()
			Tab44.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab44)
		case 45:
			Tab45 = widget.NewMultiLineEntry()
			Tab45.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab45)
		case 46:
			Tab46 = widget.NewMultiLineEntry()
			Tab46.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab46)
		case 47:
			Tab47 = widget.NewMultiLineEntry()
			Tab47.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab47)
		case 48:
			Tab48 = widget.NewMultiLineEntry()
			Tab48.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab48)
		case 49:
			Tab49 = widget.NewMultiLineEntry()
			Tab49.Wrapping = fyne.TextWrapWord
			return container.NewTabItem(share.Bible, Tab49)
		default:
			return container.NewTabItem("More than 50 tabs are not supported!", widget.NewMultiLineEntry())
		}
	}
	BibleTabs.SetTabLocation(container.TabLocationTop)
	return container.NewBorder(nil, nil, nil, nil, BibleTabs)
}

func RunCommand(command, bibleModule string, tabs *container.DocTabs) {
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
		//display bible text
		tabs.Selected().Text = share.Reference
		i := tabs.SelectedIndex()
		switch i {
		case 0:
			Tab0.SetText(bible.Display)
		case 1:
			Tab1.SetText(bible.Display)
		case 2:
			Tab2.SetText(bible.Display)
		case 3:
			Tab3.SetText(bible.Display)
		case 4:
			Tab4.SetText(bible.Display)
		case 5:
			Tab5.SetText(bible.Display)
		case 6:
			Tab6.SetText(bible.Display)
		case 7:
			Tab7.SetText(bible.Display)
		case 8:
			Tab8.SetText(bible.Display)
		case 9:
			Tab9.SetText(bible.Display)
		case 10:
			Tab10.SetText(bible.Display)
		case 11:
			Tab11.SetText(bible.Display)
		case 12:
			Tab12.SetText(bible.Display)
		case 13:
			Tab13.SetText(bible.Display)
		case 14:
			Tab14.SetText(bible.Display)
		case 15:
			Tab15.SetText(bible.Display)
		case 16:
			Tab16.SetText(bible.Display)
		case 17:
			Tab17.SetText(bible.Display)
		case 18:
			Tab18.SetText(bible.Display)
		case 19:
			Tab19.SetText(bible.Display)
		case 20:
			Tab20.SetText(bible.Display)
		case 21:
			Tab21.SetText(bible.Display)
		case 22:
			Tab22.SetText(bible.Display)
		case 23:
			Tab23.SetText(bible.Display)
		case 24:
			Tab24.SetText(bible.Display)
		case 25:
			Tab25.SetText(bible.Display)
		case 26:
			Tab26.SetText(bible.Display)
		case 27:
			Tab27.SetText(bible.Display)
		case 28:
			Tab28.SetText(bible.Display)
		case 29:
			Tab29.SetText(bible.Display)
		case 30:
			Tab30.SetText(bible.Display)
		case 31:
			Tab31.SetText(bible.Display)
		case 32:
			Tab32.SetText(bible.Display)
		case 33:
			Tab33.SetText(bible.Display)
		case 34:
			Tab34.SetText(bible.Display)
		case 35:
			Tab35.SetText(bible.Display)
		case 36:
			Tab36.SetText(bible.Display)
		case 37:
			Tab37.SetText(bible.Display)
		case 38:
			Tab38.SetText(bible.Display)
		case 39:
			Tab39.SetText(bible.Display)
		case 40:
			Tab40.SetText(bible.Display)
		case 41:
			Tab41.SetText(bible.Display)
		case 42:
			Tab42.SetText(bible.Display)
		case 43:
			Tab43.SetText(bible.Display)
		case 44:
			Tab44.SetText(bible.Display)
		case 45:
			Tab45.SetText(bible.Display)
		case 46:
			Tab46.SetText(bible.Display)
		case 47:
			Tab47.SetText(bible.Display)
		case 48:
			Tab48.SetText(bible.Display)
		case 49:
			Tab49.SetText(bible.Display)
		}
		tabs.Refresh()
	}
}
