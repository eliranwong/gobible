package fyne

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/eliranwong/gobible/internal/parser"
	"github.com/eliranwong/gobible/internal/share"
)

var searchProgress *widget.ProgressBar

func populateSearchTabs() {

	var results1, results2, results3, results4, results5, results6, results7, results8, results9 [][]string
	var results10, results11, results12, results13, results14, results15, results16, results17, results18, results19 [][]string
	var results20, results21, results22, results23, results24, results25, results26, results27, results28, results29 [][]string
	var results30, results31, results32, results33, results34, results35, results36, results37, results38, results39 [][]string
	var results40, results41, results42, results43, results44, results45, results46, results47, results48, results49 [][]string
	var results50, results51, results52, results53, results54, results55, results56, results57, results58, results59 [][]string
	var results60, results61, results62, results63, results64, results65, results66, results67 [][]string

	searchProgress.Show()
	mainLayout.Refresh()
	progress := 0.0
	progressIncrement := 100.0 / 67.0 / 100
	searchProgress.SetValue(progress)

	//searchTabItem1.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["1"], results1[1])
	//searchTabWidget1.SetText(results1[0])
	totalCount := 0
	for i := 1; i <= 67; i++ {
		select {
		case results1 = <-share.Ch1:
			totalCount += len(results1)
		case results2 = <-share.Ch2:
			totalCount += len(results2)
		case results3 = <-share.Ch3:
			totalCount += len(results3)
		case results4 = <-share.Ch4:
			totalCount += len(results4)
		case results5 = <-share.Ch5:
			totalCount += len(results5)
		case results6 = <-share.Ch6:
			totalCount += len(results6)
		case results7 = <-share.Ch7:
			totalCount += len(results7)
		case results8 = <-share.Ch8:
			totalCount += len(results8)
		case results9 = <-share.Ch9:
			totalCount += len(results9)
		case results10 = <-share.Ch10:
			totalCount += len(results10)
		case results11 = <-share.Ch11:
			totalCount += len(results11)
		case results12 = <-share.Ch12:
			totalCount += len(results12)
		case results13 = <-share.Ch13:
			totalCount += len(results13)
		case results14 = <-share.Ch14:
			totalCount += len(results14)
		case results15 = <-share.Ch15:
			totalCount += len(results15)
		case results16 = <-share.Ch16:
			totalCount += len(results16)
		case results17 = <-share.Ch17:
			totalCount += len(results17)
		case results18 = <-share.Ch18:
			totalCount += len(results18)
		case results19 = <-share.Ch19:
			totalCount += len(results19)
		case results20 = <-share.Ch20:
			totalCount += len(results20)
		case results21 = <-share.Ch21:
			totalCount += len(results21)
		case results22 = <-share.Ch22:
			totalCount += len(results22)
		case results23 = <-share.Ch23:
			totalCount += len(results23)
		case results24 = <-share.Ch24:
			totalCount += len(results24)
		case results25 = <-share.Ch25:
			totalCount += len(results25)
		case results26 = <-share.Ch26:
			totalCount += len(results26)
		case results27 = <-share.Ch27:
			totalCount += len(results27)
		case results28 = <-share.Ch28:
			totalCount += len(results28)
		case results29 = <-share.Ch29:
			totalCount += len(results29)
		case results30 = <-share.Ch30:
			totalCount += len(results30)
		case results31 = <-share.Ch31:
			totalCount += len(results31)
		case results32 = <-share.Ch32:
			totalCount += len(results32)
		case results33 = <-share.Ch33:
			totalCount += len(results33)
		case results34 = <-share.Ch34:
			totalCount += len(results34)
		case results35 = <-share.Ch35:
			totalCount += len(results35)
		case results36 = <-share.Ch36:
			totalCount += len(results36)
		case results37 = <-share.Ch37:
			totalCount += len(results37)
		case results38 = <-share.Ch38:
			totalCount += len(results38)
		case results39 = <-share.Ch39:
			totalCount += len(results39)
		case results40 = <-share.Ch40:
			totalCount += len(results40)
		case results41 = <-share.Ch41:
			totalCount += len(results41)
		case results42 = <-share.Ch42:
			totalCount += len(results42)
		case results43 = <-share.Ch43:
			totalCount += len(results43)
		case results44 = <-share.Ch44:
			totalCount += len(results44)
		case results45 = <-share.Ch45:
			totalCount += len(results45)
		case results46 = <-share.Ch46:
			totalCount += len(results46)
		case results47 = <-share.Ch47:
			totalCount += len(results47)
		case results48 = <-share.Ch48:
			totalCount += len(results48)
		case results49 = <-share.Ch49:
			totalCount += len(results49)
		case results50 = <-share.Ch50:
			totalCount += len(results50)
		case results51 = <-share.Ch51:
			totalCount += len(results51)
		case results52 = <-share.Ch52:
			totalCount += len(results52)
		case results53 = <-share.Ch53:
			totalCount += len(results53)
		case results54 = <-share.Ch54:
			totalCount += len(results54)
		case results55 = <-share.Ch55:
			totalCount += len(results55)
		case results56 = <-share.Ch56:
			totalCount += len(results56)
		case results57 = <-share.Ch57:
			totalCount += len(results57)
		case results58 = <-share.Ch58:
			totalCount += len(results58)
		case results59 = <-share.Ch59:
			totalCount += len(results59)
		case results60 = <-share.Ch60:
			totalCount += len(results60)
		case results61 = <-share.Ch61:
			totalCount += len(results61)
		case results62 = <-share.Ch62:
			totalCount += len(results62)
		case results63 = <-share.Ch63:
			totalCount += len(results63)
		case results64 = <-share.Ch64:
			totalCount += len(results64)
		case results65 = <-share.Ch65:
			totalCount += len(results65)
		case results66 = <-share.Ch66:
			totalCount += len(results66)
		case results67 = <-share.Ch67:
			totalCount += len(results67)
		}
		progress += progressIncrement
		searchProgress.SetValue(progress)
	}
	searchProgress.SetValue(1.0)
	searchProgress.Hide()
	mainLayout.Refresh()

	// combine all results
	allResults := make([][]string, totalCount)
	var i int
	i += copy(allResults[i:], results1)
	i += copy(allResults[i:], results2)
	i += copy(allResults[i:], results3)
	i += copy(allResults[i:], results4)
	i += copy(allResults[i:], results5)
	i += copy(allResults[i:], results6)
	i += copy(allResults[i:], results7)
	i += copy(allResults[i:], results8)
	i += copy(allResults[i:], results9)
	i += copy(allResults[i:], results10)
	i += copy(allResults[i:], results11)
	i += copy(allResults[i:], results12)
	i += copy(allResults[i:], results13)
	i += copy(allResults[i:], results14)
	i += copy(allResults[i:], results15)
	i += copy(allResults[i:], results16)
	i += copy(allResults[i:], results17)
	i += copy(allResults[i:], results18)
	i += copy(allResults[i:], results19)
	i += copy(allResults[i:], results20)
	i += copy(allResults[i:], results21)
	i += copy(allResults[i:], results22)
	i += copy(allResults[i:], results23)
	i += copy(allResults[i:], results24)
	i += copy(allResults[i:], results25)
	i += copy(allResults[i:], results26)
	i += copy(allResults[i:], results27)
	i += copy(allResults[i:], results28)
	i += copy(allResults[i:], results29)
	i += copy(allResults[i:], results30)
	i += copy(allResults[i:], results31)
	i += copy(allResults[i:], results32)
	i += copy(allResults[i:], results33)
	i += copy(allResults[i:], results34)
	i += copy(allResults[i:], results35)
	i += copy(allResults[i:], results36)
	i += copy(allResults[i:], results37)
	i += copy(allResults[i:], results38)
	i += copy(allResults[i:], results39)
	i += copy(allResults[i:], results40)
	i += copy(allResults[i:], results41)
	i += copy(allResults[i:], results42)
	i += copy(allResults[i:], results43)
	i += copy(allResults[i:], results44)
	i += copy(allResults[i:], results45)
	i += copy(allResults[i:], results46)
	i += copy(allResults[i:], results47)
	i += copy(allResults[i:], results48)
	i += copy(allResults[i:], results49)
	i += copy(allResults[i:], results50)
	i += copy(allResults[i:], results51)
	i += copy(allResults[i:], results52)
	i += copy(allResults[i:], results53)
	i += copy(allResults[i:], results54)
	i += copy(allResults[i:], results55)
	i += copy(allResults[i:], results56)
	i += copy(allResults[i:], results57)
	i += copy(allResults[i:], results58)
	i += copy(allResults[i:], results59)
	i += copy(allResults[i:], results60)
	i += copy(allResults[i:], results61)
	i += copy(allResults[i:], results62)
	i += copy(allResults[i:], results63)
	i += copy(allResults[i:], results64)
	i += copy(allResults[i:], results65)
	i += copy(allResults[i:], results66)
	i += copy(allResults[i:], results67)

	resultLabelAll := fmt.Sprintf("ALL [%v]", totalCount)
	resultLabel1 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["1"], len(results1))
	resultLabel2 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["2"], len(results2))
	resultLabel3 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["3"], len(results3))
	resultLabel4 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["4"], len(results4))
	resultLabel5 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["5"], len(results5))
	resultLabel6 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["6"], len(results6))
	resultLabel7 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["7"], len(results7))
	resultLabel8 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["8"], len(results8))
	resultLabel9 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["9"], len(results9))
	resultLabel10 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["10"], len(results10))
	resultLabel11 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["11"], len(results11))
	resultLabel12 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["12"], len(results12))
	resultLabel13 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["13"], len(results13))
	resultLabel14 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["14"], len(results14))
	resultLabel15 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["15"], len(results15))
	resultLabel16 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["16"], len(results16))
	resultLabel17 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["17"], len(results17))
	resultLabel18 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["18"], len(results18))
	resultLabel19 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["19"], len(results19))
	resultLabel20 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["20"], len(results20))
	resultLabel21 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["21"], len(results21))
	resultLabel22 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["22"], len(results22))
	resultLabel23 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["23"], len(results23))
	resultLabel24 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["24"], len(results24))
	resultLabel25 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["25"], len(results25))
	resultLabel26 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["26"], len(results26))
	resultLabel27 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["27"], len(results27))
	resultLabel28 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["28"], len(results28))
	resultLabel29 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["29"], len(results29))
	resultLabel30 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["30"], len(results30))
	resultLabel31 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["31"], len(results31))
	resultLabel32 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["32"], len(results32))
	resultLabel33 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["33"], len(results33))
	resultLabel34 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["34"], len(results34))
	resultLabel35 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["35"], len(results35))
	resultLabel36 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["36"], len(results36))
	resultLabel37 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["37"], len(results37))
	resultLabel38 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["38"], len(results38))
	resultLabel39 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["39"], len(results39))
	resultLabel40 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["40"], len(results40))
	resultLabel41 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["41"], len(results41))
	resultLabel42 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["42"], len(results42))
	resultLabel43 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["43"], len(results43))
	resultLabel44 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["44"], len(results44))
	resultLabel45 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["45"], len(results45))
	resultLabel46 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["46"], len(results46))
	resultLabel47 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["47"], len(results47))
	resultLabel48 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["48"], len(results48))
	resultLabel49 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["49"], len(results49))
	resultLabel50 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["50"], len(results50))
	resultLabel51 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["51"], len(results51))
	resultLabel52 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["52"], len(results52))
	resultLabel53 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["53"], len(results53))
	resultLabel54 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["54"], len(results54))
	resultLabel55 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["55"], len(results55))
	resultLabel56 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["56"], len(results56))
	resultLabel57 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["57"], len(results57))
	resultLabel58 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["58"], len(results58))
	resultLabel59 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["59"], len(results59))
	resultLabel60 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["60"], len(results60))
	resultLabel61 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["61"], len(results61))
	resultLabel62 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["62"], len(results62))
	resultLabel63 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["63"], len(results63))
	resultLabel64 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["64"], len(results64))
	resultLabel65 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["65"], len(results65))
	resultLabel66 := fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["66"], len(results66))
	resultLabel67 := fmt.Sprintf("... [%v]", len(results67))

	/*searchTabs := container.NewAppTabs(
		container.NewTabItem(resultLabelAll, makeVerseTable(allResults)),
	)*/

	searchTabs := container.NewDocTabs(
		container.NewTabItem(resultLabelAll, makeVerseTable(allResults)),
	)
	searchTabs.CreateTab = func() *container.TabItem {
		return container.NewTabItem(resultLabelAll, makeVerseTable(allResults))
	}

	searchTabs.SetTabLocation(container.TabLocationTop)

	resultList := widget.NewList(
		func() int {
			return 68
		},
		func() fyne.CanvasObject {
			// determine the width
			return widget.NewLabel("..............................")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			var labelText string
			switch i {
			case 0:
				labelText = resultLabelAll
			case 1:
				labelText = resultLabel1
			case 2:
				labelText = resultLabel2
			case 3:
				labelText = resultLabel3
			case 4:
				labelText = resultLabel4
			case 5:
				labelText = resultLabel5
			case 6:
				labelText = resultLabel6
			case 7:
				labelText = resultLabel7
			case 8:
				labelText = resultLabel8
			case 9:
				labelText = resultLabel9
			case 10:
				labelText = resultLabel10
			case 11:
				labelText = resultLabel11
			case 12:
				labelText = resultLabel12
			case 13:
				labelText = resultLabel13
			case 14:
				labelText = resultLabel14
			case 15:
				labelText = resultLabel15
			case 16:
				labelText = resultLabel16
			case 17:
				labelText = resultLabel17
			case 18:
				labelText = resultLabel18
			case 19:
				labelText = resultLabel19
			case 20:
				labelText = resultLabel20
			case 21:
				labelText = resultLabel21
			case 22:
				labelText = resultLabel22
			case 23:
				labelText = resultLabel23
			case 24:
				labelText = resultLabel24
			case 25:
				labelText = resultLabel25
			case 26:
				labelText = resultLabel26
			case 27:
				labelText = resultLabel27
			case 28:
				labelText = resultLabel28
			case 29:
				labelText = resultLabel29
			case 30:
				labelText = resultLabel30
			case 31:
				labelText = resultLabel31
			case 32:
				labelText = resultLabel32
			case 33:
				labelText = resultLabel33
			case 34:
				labelText = resultLabel34
			case 35:
				labelText = resultLabel35
			case 36:
				labelText = resultLabel36
			case 37:
				labelText = resultLabel37
			case 38:
				labelText = resultLabel38
			case 39:
				labelText = resultLabel39
			case 40:
				labelText = resultLabel40
			case 41:
				labelText = resultLabel41
			case 42:
				labelText = resultLabel42
			case 43:
				labelText = resultLabel43
			case 44:
				labelText = resultLabel44
			case 45:
				labelText = resultLabel45
			case 46:
				labelText = resultLabel46
			case 47:
				labelText = resultLabel47
			case 48:
				labelText = resultLabel48
			case 49:
				labelText = resultLabel49
			case 50:
				labelText = resultLabel50
			case 51:
				labelText = resultLabel51
			case 52:
				labelText = resultLabel52
			case 53:
				labelText = resultLabel53
			case 54:
				labelText = resultLabel54
			case 55:
				labelText = resultLabel55
			case 56:
				labelText = resultLabel56
			case 57:
				labelText = resultLabel57
			case 58:
				labelText = resultLabel58
			case 59:
				labelText = resultLabel59
			case 60:
				labelText = resultLabel60
			case 61:
				labelText = resultLabel61
			case 62:
				labelText = resultLabel62
			case 63:
				labelText = resultLabel63
			case 64:
				labelText = resultLabel64
			case 65:
				labelText = resultLabel65
			case 66:
				labelText = resultLabel66
			case 67:
				labelText = resultLabel67
			}
			o.(*widget.Label).SetText(labelText)
		})

	resultList.OnSelected = func(id widget.ListItemID) {
		switch id {
		case 0:
			searchTabs.Append(container.NewTabItem(resultLabelAll, makeVerseTable(allResults)))
		case 1:
			searchTabs.Append(container.NewTabItem(resultLabel1, makeVerseTable(results1)))
		case 2:
			searchTabs.Append(container.NewTabItem(resultLabel2, makeVerseTable(results2)))
		case 3:
			searchTabs.Append(container.NewTabItem(resultLabel3, makeVerseTable(results3)))
		case 4:
			searchTabs.Append(container.NewTabItem(resultLabel4, makeVerseTable(results4)))
		case 5:
			searchTabs.Append(container.NewTabItem(resultLabel5, makeVerseTable(results5)))
		case 6:
			searchTabs.Append(container.NewTabItem(resultLabel6, makeVerseTable(results6)))
		case 7:
			searchTabs.Append(container.NewTabItem(resultLabel7, makeVerseTable(results7)))
		case 8:
			searchTabs.Append(container.NewTabItem(resultLabel8, makeVerseTable(results8)))
		case 9:
			searchTabs.Append(container.NewTabItem(resultLabel9, makeVerseTable(results9)))
		case 10:
			searchTabs.Append(container.NewTabItem(resultLabel10, makeVerseTable(results10)))
		case 11:
			searchTabs.Append(container.NewTabItem(resultLabel11, makeVerseTable(results11)))
		case 12:
			searchTabs.Append(container.NewTabItem(resultLabel12, makeVerseTable(results12)))
		case 13:
			searchTabs.Append(container.NewTabItem(resultLabel13, makeVerseTable(results13)))
		case 14:
			searchTabs.Append(container.NewTabItem(resultLabel14, makeVerseTable(results14)))
		case 15:
			searchTabs.Append(container.NewTabItem(resultLabel15, makeVerseTable(results15)))
		case 16:
			searchTabs.Append(container.NewTabItem(resultLabel16, makeVerseTable(results16)))
		case 17:
			searchTabs.Append(container.NewTabItem(resultLabel17, makeVerseTable(results17)))
		case 18:
			searchTabs.Append(container.NewTabItem(resultLabel18, makeVerseTable(results18)))
		case 19:
			searchTabs.Append(container.NewTabItem(resultLabel19, makeVerseTable(results19)))
		case 20:
			searchTabs.Append(container.NewTabItem(resultLabel20, makeVerseTable(results20)))
		case 21:
			searchTabs.Append(container.NewTabItem(resultLabel21, makeVerseTable(results21)))
		case 22:
			searchTabs.Append(container.NewTabItem(resultLabel22, makeVerseTable(results22)))
		case 23:
			searchTabs.Append(container.NewTabItem(resultLabel23, makeVerseTable(results23)))
		case 24:
			searchTabs.Append(container.NewTabItem(resultLabel24, makeVerseTable(results24)))
		case 25:
			searchTabs.Append(container.NewTabItem(resultLabel25, makeVerseTable(results25)))
		case 26:
			searchTabs.Append(container.NewTabItem(resultLabel26, makeVerseTable(results26)))
		case 27:
			searchTabs.Append(container.NewTabItem(resultLabel27, makeVerseTable(results27)))
		case 28:
			searchTabs.Append(container.NewTabItem(resultLabel28, makeVerseTable(results28)))
		case 29:
			searchTabs.Append(container.NewTabItem(resultLabel29, makeVerseTable(results29)))
		case 30:
			searchTabs.Append(container.NewTabItem(resultLabel30, makeVerseTable(results30)))
		case 31:
			searchTabs.Append(container.NewTabItem(resultLabel31, makeVerseTable(results31)))
		case 32:
			searchTabs.Append(container.NewTabItem(resultLabel32, makeVerseTable(results32)))
		case 33:
			searchTabs.Append(container.NewTabItem(resultLabel33, makeVerseTable(results33)))
		case 34:
			searchTabs.Append(container.NewTabItem(resultLabel34, makeVerseTable(results34)))
		case 35:
			searchTabs.Append(container.NewTabItem(resultLabel35, makeVerseTable(results35)))
		case 36:
			searchTabs.Append(container.NewTabItem(resultLabel36, makeVerseTable(results36)))
		case 37:
			searchTabs.Append(container.NewTabItem(resultLabel37, makeVerseTable(results37)))
		case 38:
			searchTabs.Append(container.NewTabItem(resultLabel38, makeVerseTable(results38)))
		case 39:
			searchTabs.Append(container.NewTabItem(resultLabel39, makeVerseTable(results39)))
		case 40:
			searchTabs.Append(container.NewTabItem(resultLabel40, makeVerseTable(results40)))
		case 41:
			searchTabs.Append(container.NewTabItem(resultLabel41, makeVerseTable(results41)))
		case 42:
			searchTabs.Append(container.NewTabItem(resultLabel42, makeVerseTable(results42)))
		case 43:
			searchTabs.Append(container.NewTabItem(resultLabel43, makeVerseTable(results43)))
		case 44:
			searchTabs.Append(container.NewTabItem(resultLabel44, makeVerseTable(results44)))
		case 45:
			searchTabs.Append(container.NewTabItem(resultLabel45, makeVerseTable(results45)))
		case 46:
			searchTabs.Append(container.NewTabItem(resultLabel46, makeVerseTable(results46)))
		case 47:
			searchTabs.Append(container.NewTabItem(resultLabel47, makeVerseTable(results47)))
		case 48:
			searchTabs.Append(container.NewTabItem(resultLabel48, makeVerseTable(results48)))
		case 49:
			searchTabs.Append(container.NewTabItem(resultLabel49, makeVerseTable(results49)))
		case 50:
			searchTabs.Append(container.NewTabItem(resultLabel50, makeVerseTable(results50)))
		case 51:
			searchTabs.Append(container.NewTabItem(resultLabel51, makeVerseTable(results51)))
		case 52:
			searchTabs.Append(container.NewTabItem(resultLabel52, makeVerseTable(results52)))
		case 53:
			searchTabs.Append(container.NewTabItem(resultLabel53, makeVerseTable(results53)))
		case 54:
			searchTabs.Append(container.NewTabItem(resultLabel54, makeVerseTable(results54)))
		case 55:
			searchTabs.Append(container.NewTabItem(resultLabel55, makeVerseTable(results55)))
		case 56:
			searchTabs.Append(container.NewTabItem(resultLabel56, makeVerseTable(results56)))
		case 57:
			searchTabs.Append(container.NewTabItem(resultLabel57, makeVerseTable(results57)))
		case 58:
			searchTabs.Append(container.NewTabItem(resultLabel58, makeVerseTable(results58)))
		case 59:
			searchTabs.Append(container.NewTabItem(resultLabel59, makeVerseTable(results59)))
		case 60:
			searchTabs.Append(container.NewTabItem(resultLabel60, makeVerseTable(results60)))
		case 61:
			searchTabs.Append(container.NewTabItem(resultLabel61, makeVerseTable(results61)))
		case 62:
			searchTabs.Append(container.NewTabItem(resultLabel62, makeVerseTable(results62)))
		case 63:
			searchTabs.Append(container.NewTabItem(resultLabel63, makeVerseTable(results63)))
		case 64:
			searchTabs.Append(container.NewTabItem(resultLabel64, makeVerseTable(results64)))
		case 65:
			searchTabs.Append(container.NewTabItem(resultLabel65, makeVerseTable(results65)))
		case 66:
			searchTabs.Append(container.NewTabItem(resultLabel66, makeVerseTable(results66)))
		case 67:
			searchTabs.Append(container.NewTabItem(resultLabel67, makeVerseTable(results67)))
		}
		searchTabs.SelectIndex(len(searchTabs.Items) - 1)
	}

	searchTabsContainer := container.NewBorder(nil, nil, resultList, nil, searchTabs)

	// make window
	w := fyne.CurrentApp().NewWindow("Search Results")
	w.Resize(fyne.NewSize(800, 600))
	w.SetContent(searchTabsContainer)
	w.Show()
	//allResultsTable.ScrollTo(widget.TableCellID{Row: 0, Col: 0})
	//searchTabs.SelectIndex(0)
}

func makeVerseTable(data [][]string) *widget.Table {
	t := widget.NewTable(
		func() (int, int) { return len(data), 3 },
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			label.SetText(data[id.Row][id.Col])
		})
	t.OnSelected = func(id widget.TableCellID) {
		// prevent uncontrolled scrolling
		t.ScrollToLeading()
		share.Bible = data[id.Row][0]
		bibleSelect.PlaceHolder = share.Bible
		runText(data[id.Row][1])
	}
	t.SetColumnWidth(0, 80)
	t.SetColumnWidth(1, 100)
	// workaround to enable horizontal scrolling
	t.SetColumnWidth(2, 10000)
	// workaround to update table content
	t.ScrollToBottom()
	return t
}

func runText(text string) {
	references := parser.ExtractAllReferences(text, false)
	if !(len(references) == 0) {
		RunCommand(text, share.Bible, bibleTabs)
	}
}
