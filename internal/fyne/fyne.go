package fyne

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/eliranwong/gobible/internal/bible"
	"github.com/eliranwong/gobible/internal/check"
	"github.com/eliranwong/gobible/internal/parser"
	"github.com/eliranwong/gobible/internal/share"
)

var mainWindow fyne.Window
var mainLayout *fyne.Container
var bibleTabs *container.DocTabs
var bibleSelect *widget.SelectEntry
var bibleTab0, bibleTab1, bibleTab2, bibleTab3, bibleTab4, bibleTab5, bibleTab6, bibleTab7, bibleTab8, bibleTab9 *widget.Entry
var bibleTab10, bibleTab11, bibleTab12, bibleTab13, bibleTab14, bibleTab15, bibleTab16, bibleTab17, bibleTab18, bibleTab19 *widget.Entry
var bibleTab20, bibleTab21, bibleTab22, bibleTab23, bibleTab24, bibleTab25, bibleTab26, bibleTab27, bibleTab28, bibleTab29 *widget.Entry
var bibleTab30, bibleTab31, bibleTab32, bibleTab33, bibleTab34, bibleTab35, bibleTab36, bibleTab37, bibleTab38, bibleTab39 *widget.Entry
var bibleTab40, bibleTab41, bibleTab42, bibleTab43, bibleTab44, bibleTab45, bibleTab46, bibleTab47, bibleTab48, bibleTab49 *widget.Entry

func Fyne() {
	makeMainWindow()
	setUpUI()
	mainWindow.SetMaster()
	mainWindow.ShowAndRun()
}

func makeMainWindow() {
	share.SetupData()

	os.Setenv("FYNE_FONT", filepath.Join(share.Data, filepath.FromSlash("fonts/fonts.ttf")))
	// set appication size with FYNE_SCALE
	// read https://developer.fyne.io/architecture/scaling
	os.Setenv("FYNE_SCALE", "1.2")

	gobible := app.NewWithID("app.gobible")
	initPreferences()

	mainWindow = gobible.NewWindow("Go Bible")
	mainWindow.Resize(fyne.NewSize(1024, 768))
}

func setUpUI() {
	// tabs for displaying bible text
	bibleTabsContainer := makeDocTabsTab()
	// text entry and drowndown menu for bible selection
	bibleSelect = widget.NewSelectEntry(share.Bibles)
	bibleSelect.PlaceHolder = share.Bible
	bibleSelect.OnChanged = func(s string) {
		if !(s == "") {
			filePath := filepath.Join(share.Data, filepath.FromSlash(fmt.Sprintf("bibles/%v.bible", s)))
			if check.FileExists(filePath) {
				share.Bible = s
				bibleSelect.PlaceHolder = share.Bible
				RunCommand(share.Reference, share.Bible, bibleTabs)
			}
		}
	}
	bibleSelect.OnSubmitted = func(s string) {
		if !(s == "") {
			share.Bible = s
			RunCommand(share.Reference, share.Bible, bibleTabs)
		}
	}
	// bible passage selection tree
	chapters := makeBcvTree(true)
	reference := ""
	chapters.OnBranchOpened = func(id string) {
		reference = fmt.Sprintf(`%v %v`, id, 1)
	}
	chapters.OnSelected = func(id string) {
		_, err := strconv.Atoi(id[len(id)-1:])
		if err == nil {
			reference = id
		} else {
			reference = fmt.Sprintf(`%v %v`, id, 1)
		}
		RunCommand(reference, share.Bible, bibleTabs)
	}
	// layout putting together bible navigation elements
	bibleNavigator := container.NewBorder(bibleSelect, nil, nil, nil, chapters)
	bibleNavigator.Hide()
	// bible layout
	bibleLayout := container.NewBorder(nil, nil, bibleNavigator, nil, bibleTabsContainer)

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
	// previous and next chapter buttons
	previousChapter := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		ref := bible.GetPreviousChapterRef()
		RunCommand(ref, share.Bible, bibleTabs)
	})
	nextChapter := widget.NewButtonWithIcon("", theme.NavigateNextIcon(), func() {
		ref := bible.GetNextChapterRef()
		RunCommand(ref, share.Bible, bibleTabs)
	})
	// text entry for command
	command := widget.NewEntry()
	command.SetPlaceHolder("Enter bible reference or search item here ...")
	command.OnSubmitted = func(s string) {
		RunCommand(s, share.Bible, bibleTabs)
	}
	// top right buttons
	/* sub-menu example
	shareItem := fyne.NewMenuItem("Share via", nil)
	shareItem.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Twitter", func() { fmt.Println("context menu Share->Twitter") }),
		fyne.NewMenuItem("Reddit", func() { fmt.Println("context menu Share->Reddit") }),
	)*/
	featureMenuButton := newContextMenuButton("", theme.ContentAddIcon(), fyne.NewMenu("",
		fyne.NewMenuItem("Compare Chapter", func() { makeComparisonWindow(false, false) }),
		fyne.NewMenuItem("Compare Verse", func() { makeComparisonWindow(true, false) }),
		fyne.NewMenuItem("Cross-reference", func() { makeComparisonWindow(true, true) }),
		//fyne.NewMenuItem("Terminal", newTerminalWindow),
	))
	settingButton := newContextMenuButton("", theme.SettingsIcon(), fyne.NewMenu("",
		fyne.NewMenuItem("Toggle Theme", toggleFyneTheme),
		fyne.NewMenuItem("Quit", fyne.CurrentApp().Quit),
	))

	// display search result
	//searchDisplayArea = makeSearchTabs()
	searchProgress = widget.NewProgressBar()
	searchProgress.Hide()

	topLeftButton := container.NewHBox(previousChapter, showHideBibleNavigator, nextChapter)
	topRightButton := container.NewHBox(featureMenuButton, settingButton)
	mainTop := container.NewBorder(nil, nil, topLeftButton, topRightButton, command)
	//mainCentre := container.NewHSplit(bibleLayout, searchDisplayArea)
	mainLayout = container.NewBorder(mainTop, searchProgress, nil, nil, bibleLayout)
	mainWindow.SetContent(mainLayout)

	startupCommand := share.Reference
	command.Text = startupCommand
	RunCommand(startupCommand, share.Bible, bibleTabs)
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

func makeBcvTree(verses bool) *widget.Tree {
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
			if verses {
				vTotal = parser.Verses[b][c]
				for v := 1; v <= vTotal; v++ {
					data[chapterStr] = append(data[chapterStr], fmt.Sprintf(`%v %v:%v`, abb, cStr, v))
				}
			}
		}
	}

	tree := widget.NewTreeWithStrings(data)
	/*tree.OnUnselected = func(id string) {
		//
	}*/
	//tree.OpenBranch("John")
	return tree
}

/*
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
}*/

func makeMultiLineEntry() *widget.Entry {
	entry := widget.NewMultiLineEntry()
	entry.Wrapping = fyne.TextWrapWord
	return entry
}

func makeDocTabsTab() fyne.CanvasObject {
	bibleTab0 = makeMultiLineEntry()
	bibleTabs = container.NewDocTabs(
		container.NewTabItem(share.Bible, bibleTab0),
	)
	bibleTabs.CreateTab = func() *container.TabItem {
		i := len(bibleTabs.Items)
		switch i {
		case 0:
			bibleTab0 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab0)
		case 1:
			bibleTab1 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab1)
		case 2:
			bibleTab2 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab2)
		case 3:
			bibleTab3 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab3)
		case 4:
			bibleTab4 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab4)
		case 5:
			bibleTab5 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab5)
		case 6:
			bibleTab6 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab6)
		case 7:
			bibleTab7 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab7)
		case 8:
			bibleTab8 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab8)
		case 9:
			bibleTab9 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab9)
		case 10:
			bibleTab10 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab10)
		case 11:
			bibleTab11 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab11)
		case 12:
			bibleTab12 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab12)
		case 13:
			bibleTab13 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab13)
		case 14:
			bibleTab14 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab14)
		case 15:
			bibleTab15 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab15)
		case 16:
			bibleTab16 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab16)
		case 17:
			bibleTab17 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab17)
		case 18:
			bibleTab18 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab18)
		case 19:
			bibleTab19 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab19)
		case 20:
			bibleTab20 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab20)
		case 21:
			bibleTab21 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab21)
		case 22:
			bibleTab22 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab22)
		case 23:
			bibleTab23 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab23)
		case 24:
			bibleTab24 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab24)
		case 25:
			bibleTab25 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab25)
		case 26:
			bibleTab26 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab26)
		case 27:
			bibleTab27 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab27)
		case 28:
			bibleTab28 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab28)
		case 29:
			bibleTab29 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab29)
		case 30:
			bibleTab30 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab30)
		case 31:
			bibleTab31 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab31)
		case 32:
			bibleTab32 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab32)
		case 33:
			bibleTab33 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab33)
		case 34:
			bibleTab34 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab34)
		case 35:
			bibleTab35 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab35)
		case 36:
			bibleTab36 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab36)
		case 37:
			bibleTab37 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab37)
		case 38:
			bibleTab38 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab38)
		case 39:
			bibleTab39 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab39)
		case 40:
			bibleTab40 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab40)
		case 41:
			bibleTab41 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab41)
		case 42:
			bibleTab42 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab42)
		case 43:
			bibleTab43 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab43)
		case 44:
			bibleTab44 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab44)
		case 45:
			bibleTab45 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab45)
		case 46:
			bibleTab46 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab46)
		case 47:
			bibleTab47 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab47)
		case 48:
			bibleTab48 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab48)
		case 49:
			bibleTab49 = makeMultiLineEntry()
			return container.NewTabItem(share.Bible, bibleTab49)
		default:
			return container.NewTabItem("More than 50 tabs are not supported!", makeMultiLineEntry())
		}
	}
	bibleTabs.SetTabLocation(container.TabLocationTop)
	return container.NewBorder(nil, nil, nil, nil, bibleTabs)
}

func RunCommand(command, bibleModule string, tabs *container.DocTabs) {
	if !(strings.TrimSpace(command) == "") {
		// parse bible reference
		references := parser.ExtractAllReferences(command, false)
		// search bible when there is no valid bible reference
		if len(references) == 0 {
			go bible.AndSearch(bibleModule, command)
			showSearchResults()
		} else {
			// reset bible chapter for display
			bible.Chapter.Reset()

			bible.Read(bibleModule, references)

			//display bible text
			tabs.Selected().Text = share.Reference
			i := tabs.SelectedIndex()
			switch i {
			case 0:
				bibleTab0.SetText(bible.Chapter.String())
			case 1:
				bibleTab1.SetText(bible.Chapter.String())
			case 2:
				bibleTab2.SetText(bible.Chapter.String())
			case 3:
				bibleTab3.SetText(bible.Chapter.String())
			case 4:
				bibleTab4.SetText(bible.Chapter.String())
			case 5:
				bibleTab5.SetText(bible.Chapter.String())
			case 6:
				bibleTab6.SetText(bible.Chapter.String())
			case 7:
				bibleTab7.SetText(bible.Chapter.String())
			case 8:
				bibleTab8.SetText(bible.Chapter.String())
			case 9:
				bibleTab9.SetText(bible.Chapter.String())
			case 10:
				bibleTab10.SetText(bible.Chapter.String())
			case 11:
				bibleTab11.SetText(bible.Chapter.String())
			case 12:
				bibleTab12.SetText(bible.Chapter.String())
			case 13:
				bibleTab13.SetText(bible.Chapter.String())
			case 14:
				bibleTab14.SetText(bible.Chapter.String())
			case 15:
				bibleTab15.SetText(bible.Chapter.String())
			case 16:
				bibleTab16.SetText(bible.Chapter.String())
			case 17:
				bibleTab17.SetText(bible.Chapter.String())
			case 18:
				bibleTab18.SetText(bible.Chapter.String())
			case 19:
				bibleTab19.SetText(bible.Chapter.String())
			case 20:
				bibleTab20.SetText(bible.Chapter.String())
			case 21:
				bibleTab21.SetText(bible.Chapter.String())
			case 22:
				bibleTab22.SetText(bible.Chapter.String())
			case 23:
				bibleTab23.SetText(bible.Chapter.String())
			case 24:
				bibleTab24.SetText(bible.Chapter.String())
			case 25:
				bibleTab25.SetText(bible.Chapter.String())
			case 26:
				bibleTab26.SetText(bible.Chapter.String())
			case 27:
				bibleTab27.SetText(bible.Chapter.String())
			case 28:
				bibleTab28.SetText(bible.Chapter.String())
			case 29:
				bibleTab29.SetText(bible.Chapter.String())
			case 30:
				bibleTab30.SetText(bible.Chapter.String())
			case 31:
				bibleTab31.SetText(bible.Chapter.String())
			case 32:
				bibleTab32.SetText(bible.Chapter.String())
			case 33:
				bibleTab33.SetText(bible.Chapter.String())
			case 34:
				bibleTab34.SetText(bible.Chapter.String())
			case 35:
				bibleTab35.SetText(bible.Chapter.String())
			case 36:
				bibleTab36.SetText(bible.Chapter.String())
			case 37:
				bibleTab37.SetText(bible.Chapter.String())
			case 38:
				bibleTab38.SetText(bible.Chapter.String())
			case 39:
				bibleTab39.SetText(bible.Chapter.String())
			case 40:
				bibleTab40.SetText(bible.Chapter.String())
			case 41:
				bibleTab41.SetText(bible.Chapter.String())
			case 42:
				bibleTab42.SetText(bible.Chapter.String())
			case 43:
				bibleTab43.SetText(bible.Chapter.String())
			case 44:
				bibleTab44.SetText(bible.Chapter.String())
			case 45:
				bibleTab45.SetText(bible.Chapter.String())
			case 46:
				bibleTab46.SetText(bible.Chapter.String())
			case 47:
				bibleTab47.SetText(bible.Chapter.String())
			case 48:
				bibleTab48.SetText(bible.Chapter.String())
			case 49:
				bibleTab49.SetText(bible.Chapter.String())
			}
			tabs.Refresh()
			go savePreferences()

		}

	}
}
