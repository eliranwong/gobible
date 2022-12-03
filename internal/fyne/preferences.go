package fyne

import (
	"sort"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"github.com/eliranwong/gobible/internal/share"
)

func initPreferences() {
	preferences := fyne.CurrentApp().Preferences()
	share.Bible = preferences.StringWithFallback("bible", "NET")
	selectedBibles := preferences.StringWithFallback("selectedBibles", "KJV_NET")
	if selectedBibles == "" {
		selectedBibles = "KJV_NET"
	}
	share.SelectedBibles = strings.Split(selectedBibles, "_")
	share.BookName = preferences.StringWithFallback("bookName", "John")
	share.BookAbb = preferences.StringWithFallback("bookAbb", "John")
	share.Reference = preferences.StringWithFallback("reference", "John 3:16")
	share.Book = preferences.IntWithFallback("book", 43)
	share.Chapter = preferences.IntWithFallback("chapter", 3)
	share.Verse = preferences.IntWithFallback("verse", 16)

	// temporary solution for theme setting
	// theme.DarkTheme is deprecated: This method ignores user preference and should not be used, it will be removed in v3.0.
	share.FyneTheme = preferences.StringWithFallback("fyneTheme", "dark")
	if share.FyneTheme == "dark" {
		fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())
	} else {
		fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
	}
	//share.Check()
}

func saveSelectedBibles() {
	if len(share.SelectedBibles) == 0 {
		share.SelectedBibles = []string{"KJV", "NET"}
	}
	sort.Strings(share.SelectedBibles)
	selectedBibles := strings.Join(share.SelectedBibles, "_")
	fyne.CurrentApp().Preferences().SetString("selectedBibles", selectedBibles)
}

func toggleFyneTheme() {
	if share.FyneTheme == "dark" {
		share.FyneTheme = "light"
		fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
	} else {
		share.FyneTheme = "dark"
		fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())
	}
	go fyne.CurrentApp().Preferences().SetString("fyneTheme", share.FyneTheme)
}

func savePreferences() {
	preferences := fyne.CurrentApp().Preferences()
	//preferences.SetString("fyneTheme", share.FyneTheme)
	preferences.SetString("bible", share.Bible)
	preferences.SetString("bookName", share.BookName)
	preferences.SetString("bookAbb", share.BookAbb)
	preferences.SetString("reference", share.Reference)
	preferences.SetInt("book", share.Book)
	preferences.SetInt("chapter", share.Chapter)
	preferences.SetInt("verse", share.Verse)
	//share.Check()
}
