package fyne

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/eliranwong/gobible/internal/bible"
	"github.com/eliranwong/gobible/internal/share"
)

func makeComparisonTable(reference string, verse bool, xrefs bool) *widget.Table {
	if xrefs {
		return makeVerseTable(bible.GetXrefs(reference))
	} else if verse {
		return makeVerseTable(bible.CompareVerse(reference))
	}
	return makeVerseTable(bible.CompareChapter(reference))
}

func makeComparisonWindow(verse bool, xrefs bool) {
	reference := share.Reference

	tabs := container.NewDocTabs(
		container.NewTabItem(share.Reference, makeComparisonTable(share.Reference, verse, xrefs)),
	)
	tabs.CreateTab = func() *container.TabItem {
		return container.NewTabItem(share.Reference, makeComparisonTable(share.Reference, verse, xrefs))
	}
	tabs.SetTabLocation(container.TabLocationTop)

	// bible version selection
	checkGroup := widget.NewCheckGroup(share.Bibles, func(s []string) { share.SelectedBibles = s; go saveSelectedBibles() })
	checkGroup.SetSelected(share.SelectedBibles)
	selectAll := widget.NewButton("All", func() { checkGroup.SetSelected(share.Bibles) })
	unselectAll := widget.NewButton("None", func() { checkGroup.SetSelected([]string{}) })
	vert := container.NewVScroll(container.NewVBox(selectAll, unselectAll, checkGroup))

	// bible passage selection tree
	chapters := makeBcvTree(verse)
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
		share.Reference = reference
		tabs.Append(container.NewTabItem(share.Reference, makeComparisonTable(share.Reference, verse, xrefs)))
		tabs.SelectIndex(len(tabs.Items) - 1)
	}

	content := container.NewBorder(nil, nil, chapters, vert, tabs)

	// make window
	w := fyne.CurrentApp().NewWindow("Comparison")
	w.Resize(fyne.NewSize(800, 600))
	w.SetContent(content)
	w.Show()
}
