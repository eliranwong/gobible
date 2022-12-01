package fyne

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/eliranwong/gobible/internal/parser"
	"github.com/eliranwong/gobible/internal/share"
)

var searchDisplayArea fyne.CanvasObject
var searchProgress *widget.ProgressBar
var searchTabItem1, searchTabItem2, searchTabItem3, searchTabItem4, searchTabItem5, searchTabItem6, searchTabItem7, searchTabItem8, searchTabItem9 *container.TabItem
var searchTabItem10, searchTabItem11, searchTabItem12, searchTabItem13, searchTabItem14, searchTabItem15, searchTabItem16, searchTabItem17, searchTabItem18, searchTabItem19 *container.TabItem
var searchTabItem20, searchTabItem21, searchTabItem22, searchTabItem23, searchTabItem24, searchTabItem25, searchTabItem26, searchTabItem27, searchTabItem28, searchTabItem29 *container.TabItem
var searchTabItem30, searchTabItem31, searchTabItem32, searchTabItem33, searchTabItem34, searchTabItem35, searchTabItem36, searchTabItem37, searchTabItem38, searchTabItem39 *container.TabItem
var searchTabItem40, searchTabItem41, searchTabItem42, searchTabItem43, searchTabItem44, searchTabItem45, searchTabItem46, searchTabItem47, searchTabItem48, searchTabItem49 *container.TabItem
var searchTabItem50, searchTabItem51, searchTabItem52, searchTabItem53, searchTabItem54, searchTabItem55, searchTabItem56, searchTabItem57, searchTabItem58, searchTabItem59 *container.TabItem
var searchTabItem60, searchTabItem61, searchTabItem62, searchTabItem63, searchTabItem64, searchTabItem65, searchTabItem66, searchTabItem67 *container.TabItem
var searchTabWidget1, searchTabWidget2, searchTabWidget3, searchTabWidget4, searchTabWidget5, searchTabWidget6, searchTabWidget7, searchTabWidget8, searchTabWidget9 *widget.Entry
var searchTabWidget10, searchTabWidget11, searchTabWidget12, searchTabWidget13, searchTabWidget14, searchTabWidget15, searchTabWidget16, searchTabWidget17, searchTabWidget18, searchTabWidget19 *widget.Entry
var searchTabWidget20, searchTabWidget21, searchTabWidget22, searchTabWidget23, searchTabWidget24, searchTabWidget25, searchTabWidget26, searchTabWidget27, searchTabWidget28, searchTabWidget29 *widget.Entry
var searchTabWidget30, searchTabWidget31, searchTabWidget32, searchTabWidget33, searchTabWidget34, searchTabWidget35, searchTabWidget36, searchTabWidget37, searchTabWidget38, searchTabWidget39 *widget.Entry
var searchTabWidget40, searchTabWidget41, searchTabWidget42, searchTabWidget43, searchTabWidget44, searchTabWidget45, searchTabWidget46, searchTabWidget47, searchTabWidget48, searchTabWidget49 *widget.Entry
var searchTabWidget50, searchTabWidget51, searchTabWidget52, searchTabWidget53, searchTabWidget54, searchTabWidget55, searchTabWidget56, searchTabWidget57, searchTabWidget58, searchTabWidget59 *widget.Entry
var searchTabWidget60, searchTabWidget61, searchTabWidget62, searchTabWidget63, searchTabWidget64, searchTabWidget65, searchTabWidget66, searchTabWidget67 *widget.Entry

func makeSearchTabs() fyne.CanvasObject {

	searchTabWidget1 = makeMultiLineEntry()
	searchTabWidget2 = makeMultiLineEntry()
	searchTabWidget3 = makeMultiLineEntry()
	searchTabWidget4 = makeMultiLineEntry()
	searchTabWidget5 = makeMultiLineEntry()
	searchTabWidget6 = makeMultiLineEntry()
	searchTabWidget7 = makeMultiLineEntry()
	searchTabWidget8 = makeMultiLineEntry()
	searchTabWidget9 = makeMultiLineEntry()
	searchTabWidget10 = makeMultiLineEntry()
	searchTabWidget11 = makeMultiLineEntry()
	searchTabWidget12 = makeMultiLineEntry()
	searchTabWidget13 = makeMultiLineEntry()
	searchTabWidget14 = makeMultiLineEntry()
	searchTabWidget15 = makeMultiLineEntry()
	searchTabWidget16 = makeMultiLineEntry()
	searchTabWidget17 = makeMultiLineEntry()
	searchTabWidget18 = makeMultiLineEntry()
	searchTabWidget19 = makeMultiLineEntry()
	searchTabWidget20 = makeMultiLineEntry()
	searchTabWidget21 = makeMultiLineEntry()
	searchTabWidget22 = makeMultiLineEntry()
	searchTabWidget23 = makeMultiLineEntry()
	searchTabWidget24 = makeMultiLineEntry()
	searchTabWidget25 = makeMultiLineEntry()
	searchTabWidget26 = makeMultiLineEntry()
	searchTabWidget27 = makeMultiLineEntry()
	searchTabWidget28 = makeMultiLineEntry()
	searchTabWidget29 = makeMultiLineEntry()
	searchTabWidget30 = makeMultiLineEntry()
	searchTabWidget31 = makeMultiLineEntry()
	searchTabWidget32 = makeMultiLineEntry()
	searchTabWidget33 = makeMultiLineEntry()
	searchTabWidget34 = makeMultiLineEntry()
	searchTabWidget35 = makeMultiLineEntry()
	searchTabWidget36 = makeMultiLineEntry()
	searchTabWidget37 = makeMultiLineEntry()
	searchTabWidget38 = makeMultiLineEntry()
	searchTabWidget39 = makeMultiLineEntry()
	searchTabWidget40 = makeMultiLineEntry()
	searchTabWidget41 = makeMultiLineEntry()
	searchTabWidget42 = makeMultiLineEntry()
	searchTabWidget43 = makeMultiLineEntry()
	searchTabWidget44 = makeMultiLineEntry()
	searchTabWidget45 = makeMultiLineEntry()
	searchTabWidget46 = makeMultiLineEntry()
	searchTabWidget47 = makeMultiLineEntry()
	searchTabWidget48 = makeMultiLineEntry()
	searchTabWidget49 = makeMultiLineEntry()
	searchTabWidget50 = makeMultiLineEntry()
	searchTabWidget51 = makeMultiLineEntry()
	searchTabWidget52 = makeMultiLineEntry()
	searchTabWidget53 = makeMultiLineEntry()
	searchTabWidget54 = makeMultiLineEntry()
	searchTabWidget55 = makeMultiLineEntry()
	searchTabWidget56 = makeMultiLineEntry()
	searchTabWidget57 = makeMultiLineEntry()
	searchTabWidget58 = makeMultiLineEntry()
	searchTabWidget59 = makeMultiLineEntry()
	searchTabWidget60 = makeMultiLineEntry()
	searchTabWidget61 = makeMultiLineEntry()
	searchTabWidget62 = makeMultiLineEntry()
	searchTabWidget63 = makeMultiLineEntry()
	searchTabWidget64 = makeMultiLineEntry()
	searchTabWidget65 = makeMultiLineEntry()
	searchTabWidget66 = makeMultiLineEntry()
	searchTabWidget67 = makeMultiLineEntry()

	searchTabItem1 = container.NewTabItem(parser.StandardAbbreviation["1"], searchTabWidget1)
	searchTabItem2 = container.NewTabItem(parser.StandardAbbreviation["2"], searchTabWidget2)
	searchTabItem3 = container.NewTabItem(parser.StandardAbbreviation["3"], searchTabWidget3)
	searchTabItem4 = container.NewTabItem(parser.StandardAbbreviation["4"], searchTabWidget4)
	searchTabItem5 = container.NewTabItem(parser.StandardAbbreviation["5"], searchTabWidget5)
	searchTabItem6 = container.NewTabItem(parser.StandardAbbreviation["6"], searchTabWidget6)
	searchTabItem7 = container.NewTabItem(parser.StandardAbbreviation["7"], searchTabWidget7)
	searchTabItem8 = container.NewTabItem(parser.StandardAbbreviation["8"], searchTabWidget8)
	searchTabItem9 = container.NewTabItem(parser.StandardAbbreviation["9"], searchTabWidget9)
	searchTabItem10 = container.NewTabItem(parser.StandardAbbreviation["10"], searchTabWidget10)
	searchTabItem11 = container.NewTabItem(parser.StandardAbbreviation["11"], searchTabWidget11)
	searchTabItem12 = container.NewTabItem(parser.StandardAbbreviation["12"], searchTabWidget12)
	searchTabItem13 = container.NewTabItem(parser.StandardAbbreviation["13"], searchTabWidget13)
	searchTabItem14 = container.NewTabItem(parser.StandardAbbreviation["14"], searchTabWidget14)
	searchTabItem15 = container.NewTabItem(parser.StandardAbbreviation["15"], searchTabWidget15)
	searchTabItem16 = container.NewTabItem(parser.StandardAbbreviation["16"], searchTabWidget16)
	searchTabItem17 = container.NewTabItem(parser.StandardAbbreviation["17"], searchTabWidget17)
	searchTabItem18 = container.NewTabItem(parser.StandardAbbreviation["18"], searchTabWidget18)
	searchTabItem19 = container.NewTabItem(parser.StandardAbbreviation["19"], searchTabWidget19)
	searchTabItem20 = container.NewTabItem(parser.StandardAbbreviation["20"], searchTabWidget20)
	searchTabItem21 = container.NewTabItem(parser.StandardAbbreviation["21"], searchTabWidget21)
	searchTabItem22 = container.NewTabItem(parser.StandardAbbreviation["22"], searchTabWidget22)
	searchTabItem23 = container.NewTabItem(parser.StandardAbbreviation["23"], searchTabWidget23)
	searchTabItem24 = container.NewTabItem(parser.StandardAbbreviation["24"], searchTabWidget24)
	searchTabItem25 = container.NewTabItem(parser.StandardAbbreviation["25"], searchTabWidget25)
	searchTabItem26 = container.NewTabItem(parser.StandardAbbreviation["26"], searchTabWidget26)
	searchTabItem27 = container.NewTabItem(parser.StandardAbbreviation["27"], searchTabWidget27)
	searchTabItem28 = container.NewTabItem(parser.StandardAbbreviation["28"], searchTabWidget28)
	searchTabItem29 = container.NewTabItem(parser.StandardAbbreviation["29"], searchTabWidget29)
	searchTabItem30 = container.NewTabItem(parser.StandardAbbreviation["30"], searchTabWidget30)
	searchTabItem31 = container.NewTabItem(parser.StandardAbbreviation["31"], searchTabWidget31)
	searchTabItem32 = container.NewTabItem(parser.StandardAbbreviation["32"], searchTabWidget32)
	searchTabItem33 = container.NewTabItem(parser.StandardAbbreviation["33"], searchTabWidget33)
	searchTabItem34 = container.NewTabItem(parser.StandardAbbreviation["34"], searchTabWidget34)
	searchTabItem35 = container.NewTabItem(parser.StandardAbbreviation["35"], searchTabWidget35)
	searchTabItem36 = container.NewTabItem(parser.StandardAbbreviation["36"], searchTabWidget36)
	searchTabItem37 = container.NewTabItem(parser.StandardAbbreviation["37"], searchTabWidget37)
	searchTabItem38 = container.NewTabItem(parser.StandardAbbreviation["38"], searchTabWidget38)
	searchTabItem39 = container.NewTabItem(parser.StandardAbbreviation["39"], searchTabWidget39)
	searchTabItem40 = container.NewTabItem(parser.StandardAbbreviation["40"], searchTabWidget40)
	searchTabItem41 = container.NewTabItem(parser.StandardAbbreviation["41"], searchTabWidget41)
	searchTabItem42 = container.NewTabItem(parser.StandardAbbreviation["42"], searchTabWidget42)
	searchTabItem43 = container.NewTabItem(parser.StandardAbbreviation["43"], searchTabWidget43)
	searchTabItem44 = container.NewTabItem(parser.StandardAbbreviation["44"], searchTabWidget44)
	searchTabItem45 = container.NewTabItem(parser.StandardAbbreviation["45"], searchTabWidget45)
	searchTabItem46 = container.NewTabItem(parser.StandardAbbreviation["46"], searchTabWidget46)
	searchTabItem47 = container.NewTabItem(parser.StandardAbbreviation["47"], searchTabWidget47)
	searchTabItem48 = container.NewTabItem(parser.StandardAbbreviation["48"], searchTabWidget48)
	searchTabItem49 = container.NewTabItem(parser.StandardAbbreviation["49"], searchTabWidget49)
	searchTabItem50 = container.NewTabItem(parser.StandardAbbreviation["50"], searchTabWidget50)
	searchTabItem51 = container.NewTabItem(parser.StandardAbbreviation["51"], searchTabWidget51)
	searchTabItem52 = container.NewTabItem(parser.StandardAbbreviation["52"], searchTabWidget52)
	searchTabItem53 = container.NewTabItem(parser.StandardAbbreviation["53"], searchTabWidget53)
	searchTabItem54 = container.NewTabItem(parser.StandardAbbreviation["54"], searchTabWidget54)
	searchTabItem55 = container.NewTabItem(parser.StandardAbbreviation["55"], searchTabWidget55)
	searchTabItem56 = container.NewTabItem(parser.StandardAbbreviation["56"], searchTabWidget56)
	searchTabItem57 = container.NewTabItem(parser.StandardAbbreviation["57"], searchTabWidget57)
	searchTabItem58 = container.NewTabItem(parser.StandardAbbreviation["58"], searchTabWidget58)
	searchTabItem59 = container.NewTabItem(parser.StandardAbbreviation["59"], searchTabWidget59)
	searchTabItem60 = container.NewTabItem(parser.StandardAbbreviation["60"], searchTabWidget60)
	searchTabItem61 = container.NewTabItem(parser.StandardAbbreviation["61"], searchTabWidget61)
	searchTabItem62 = container.NewTabItem(parser.StandardAbbreviation["62"], searchTabWidget62)
	searchTabItem63 = container.NewTabItem(parser.StandardAbbreviation["63"], searchTabWidget63)
	searchTabItem64 = container.NewTabItem(parser.StandardAbbreviation["64"], searchTabWidget64)
	searchTabItem65 = container.NewTabItem(parser.StandardAbbreviation["65"], searchTabWidget65)
	searchTabItem66 = container.NewTabItem(parser.StandardAbbreviation["66"], searchTabWidget66)
	searchTabItem67 = container.NewTabItem("...", searchTabWidget67)

	searchTabs := container.NewAppTabs(
		searchTabItem1,
		searchTabItem2,
		searchTabItem3,
		searchTabItem4,
		searchTabItem5,
		searchTabItem6,
		searchTabItem7,
		searchTabItem8,
		searchTabItem9,
		searchTabItem10,
		searchTabItem11,
		searchTabItem12,
		searchTabItem13,
		searchTabItem14,
		searchTabItem15,
		searchTabItem16,
		searchTabItem17,
		searchTabItem18,
		searchTabItem19,
		searchTabItem20,
		searchTabItem21,
		searchTabItem22,
		searchTabItem23,
		searchTabItem24,
		searchTabItem25,
		searchTabItem26,
		searchTabItem27,
		searchTabItem28,
		searchTabItem29,
		searchTabItem30,
		searchTabItem31,
		searchTabItem32,
		searchTabItem33,
		searchTabItem34,
		searchTabItem35,
		searchTabItem36,
		searchTabItem37,
		searchTabItem38,
		searchTabItem39,
		searchTabItem40,
		searchTabItem41,
		searchTabItem42,
		searchTabItem43,
		searchTabItem44,
		searchTabItem45,
		searchTabItem46,
		searchTabItem47,
		searchTabItem48,
		searchTabItem49,
		searchTabItem50,
		searchTabItem51,
		searchTabItem52,
		searchTabItem53,
		searchTabItem54,
		searchTabItem55,
		searchTabItem56,
		searchTabItem57,
		searchTabItem58,
		searchTabItem59,
		searchTabItem60,
		searchTabItem61,
		searchTabItem62,
		searchTabItem63,
		searchTabItem64,
		searchTabItem65,
		searchTabItem66,
		searchTabItem67,
	)
	searchTabs.SetTabLocation(container.TabLocationTop)
	searchProgress = widget.NewProgressBar()
	searchProgress.Hide()
	return container.NewBorder(nil, searchProgress, nil, nil, searchTabs)
}

func populateSearchTabs() {
	searchProgress.Show()
	searchDisplayArea.Refresh()
	progress := 0.0
	searchProgress.SetValue(progress)
	// populate tabs as soon as available
	for i := 1; i <= 67; i++ {
		select {
		case results1 := <-share.Ch1:
			searchTabItem1.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["1"], results1[1])
			searchTabWidget1.SetText(results1[0])
		case results2 := <-share.Ch2:
			searchTabItem2.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["2"], results2[1])
			searchTabWidget2.SetText(results2[0])
		case results3 := <-share.Ch3:
			searchTabItem3.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["3"], results3[1])
			searchTabWidget3.SetText(results3[0])
		case results4 := <-share.Ch4:
			searchTabItem4.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["4"], results4[1])
			searchTabWidget4.SetText(results4[0])
		case results5 := <-share.Ch5:
			searchTabItem5.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["5"], results5[1])
			searchTabWidget5.SetText(results5[0])
		case results6 := <-share.Ch6:
			searchTabItem6.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["6"], results6[1])
			searchTabWidget6.SetText(results6[0])
		case results7 := <-share.Ch7:
			searchTabItem7.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["7"], results7[1])
			searchTabWidget7.SetText(results7[0])
		case results8 := <-share.Ch8:
			searchTabItem8.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["8"], results8[1])
			searchTabWidget8.SetText(results8[0])
		case results9 := <-share.Ch9:
			searchTabItem9.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["9"], results9[1])
			searchTabWidget9.SetText(results9[0])
		case results10 := <-share.Ch10:
			searchTabItem10.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["10"], results10[1])
			searchTabWidget10.SetText(results10[0])
		case results11 := <-share.Ch11:
			searchTabItem11.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["11"], results11[1])
			searchTabWidget11.SetText(results11[0])
		case results12 := <-share.Ch12:
			searchTabItem12.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["12"], results12[1])
			searchTabWidget12.SetText(results12[0])
		case results13 := <-share.Ch13:
			searchTabItem13.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["13"], results13[1])
			searchTabWidget13.SetText(results13[0])
		case results14 := <-share.Ch14:
			searchTabItem14.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["14"], results14[1])
			searchTabWidget14.SetText(results14[0])
		case results15 := <-share.Ch15:
			searchTabItem15.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["15"], results15[1])
			searchTabWidget15.SetText(results15[0])
		case results16 := <-share.Ch16:
			searchTabItem16.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["16"], results16[1])
			searchTabWidget16.SetText(results16[0])
		case results17 := <-share.Ch17:
			searchTabItem17.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["17"], results17[1])
			searchTabWidget17.SetText(results17[0])
		case results18 := <-share.Ch18:
			searchTabItem18.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["18"], results18[1])
			searchTabWidget18.SetText(results18[0])
		case results19 := <-share.Ch19:
			searchTabItem19.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["19"], results19[1])
			searchTabWidget19.SetText(results19[0])
		case results20 := <-share.Ch20:
			searchTabItem20.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["20"], results20[1])
			searchTabWidget20.SetText(results20[0])
		case results21 := <-share.Ch21:
			searchTabItem21.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["21"], results21[1])
			searchTabWidget21.SetText(results21[0])
		case results22 := <-share.Ch22:
			searchTabItem22.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["22"], results22[1])
			searchTabWidget22.SetText(results22[0])
		case results23 := <-share.Ch23:
			searchTabItem23.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["23"], results23[1])
			searchTabWidget23.SetText(results23[0])
		case results24 := <-share.Ch24:
			searchTabItem24.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["24"], results24[1])
			searchTabWidget24.SetText(results24[0])
		case results25 := <-share.Ch25:
			searchTabItem25.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["25"], results25[1])
			searchTabWidget25.SetText(results25[0])
		case results26 := <-share.Ch26:
			searchTabItem26.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["26"], results26[1])
			searchTabWidget26.SetText(results26[0])
		case results27 := <-share.Ch27:
			searchTabItem27.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["27"], results27[1])
			searchTabWidget27.SetText(results27[0])
		case results28 := <-share.Ch28:
			searchTabItem28.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["28"], results28[1])
			searchTabWidget28.SetText(results28[0])
		case results29 := <-share.Ch29:
			searchTabItem29.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["29"], results29[1])
			searchTabWidget29.SetText(results29[0])
		case results30 := <-share.Ch30:
			searchTabItem30.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["30"], results30[1])
			searchTabWidget30.SetText(results30[0])
		case results31 := <-share.Ch31:
			searchTabItem31.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["31"], results31[1])
			searchTabWidget31.SetText(results31[0])
		case results32 := <-share.Ch32:
			searchTabItem32.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["32"], results32[1])
			searchTabWidget32.SetText(results32[0])
		case results33 := <-share.Ch33:
			searchTabItem33.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["33"], results33[1])
			searchTabWidget33.SetText(results33[0])
		case results34 := <-share.Ch34:
			searchTabItem34.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["34"], results34[1])
			searchTabWidget34.SetText(results34[0])
		case results35 := <-share.Ch35:
			searchTabItem35.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["35"], results35[1])
			searchTabWidget35.SetText(results35[0])
		case results36 := <-share.Ch36:
			searchTabItem36.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["36"], results36[1])
			searchTabWidget36.SetText(results36[0])
		case results37 := <-share.Ch37:
			searchTabItem37.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["37"], results37[1])
			searchTabWidget37.SetText(results37[0])
		case results38 := <-share.Ch38:
			searchTabItem38.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["38"], results38[1])
			searchTabWidget38.SetText(results38[0])
		case results39 := <-share.Ch39:
			searchTabItem39.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["39"], results39[1])
			searchTabWidget39.SetText(results39[0])
		case results40 := <-share.Ch40:
			searchTabItem40.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["40"], results40[1])
			searchTabWidget40.SetText(results40[0])
		case results41 := <-share.Ch41:
			searchTabItem41.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["41"], results41[1])
			searchTabWidget41.SetText(results41[0])
		case results42 := <-share.Ch42:
			searchTabItem42.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["42"], results42[1])
			searchTabWidget42.SetText(results42[0])
		case results43 := <-share.Ch43:
			searchTabItem43.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["43"], results43[1])
			searchTabWidget43.SetText(results43[0])
		case results44 := <-share.Ch44:
			searchTabItem44.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["44"], results44[1])
			searchTabWidget44.SetText(results44[0])
		case results45 := <-share.Ch45:
			searchTabItem45.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["45"], results45[1])
			searchTabWidget45.SetText(results45[0])
		case results46 := <-share.Ch46:
			searchTabItem46.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["46"], results46[1])
			searchTabWidget46.SetText(results46[0])
		case results47 := <-share.Ch47:
			searchTabItem47.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["47"], results47[1])
			searchTabWidget47.SetText(results47[0])
		case results48 := <-share.Ch48:
			searchTabItem48.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["48"], results48[1])
			searchTabWidget48.SetText(results48[0])
		case results49 := <-share.Ch49:
			searchTabItem49.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["49"], results49[1])
			searchTabWidget49.SetText(results49[0])
		case results50 := <-share.Ch50:
			searchTabItem50.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["50"], results50[1])
			searchTabWidget50.SetText(results50[0])
		case results51 := <-share.Ch51:
			searchTabItem51.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["51"], results51[1])
			searchTabWidget51.SetText(results51[0])
		case results52 := <-share.Ch52:
			searchTabItem52.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["52"], results52[1])
			searchTabWidget52.SetText(results52[0])
		case results53 := <-share.Ch53:
			searchTabItem53.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["53"], results53[1])
			searchTabWidget53.SetText(results53[0])
		case results54 := <-share.Ch54:
			searchTabItem54.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["54"], results54[1])
			searchTabWidget54.SetText(results54[0])
		case results55 := <-share.Ch55:
			searchTabItem55.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["55"], results55[1])
			searchTabWidget55.SetText(results55[0])
		case results56 := <-share.Ch56:
			searchTabItem56.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["56"], results56[1])
			searchTabWidget56.SetText(results56[0])
		case results57 := <-share.Ch57:
			searchTabItem57.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["57"], results57[1])
			searchTabWidget57.SetText(results57[0])
		case results58 := <-share.Ch58:
			searchTabItem58.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["58"], results58[1])
			searchTabWidget58.SetText(results58[0])
		case results59 := <-share.Ch59:
			searchTabItem59.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["59"], results59[1])
			searchTabWidget59.SetText(results59[0])
		case results60 := <-share.Ch60:
			searchTabItem60.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["60"], results60[1])
			searchTabWidget60.SetText(results60[0])
		case results61 := <-share.Ch61:
			searchTabItem61.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["61"], results61[1])
			searchTabWidget61.SetText(results61[0])
		case results62 := <-share.Ch62:
			searchTabItem62.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["62"], results62[1])
			searchTabWidget62.SetText(results62[0])
		case results63 := <-share.Ch63:
			searchTabItem63.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["63"], results63[1])
			searchTabWidget63.SetText(results63[0])
		case results64 := <-share.Ch64:
			searchTabItem64.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["64"], results64[1])
			searchTabWidget64.SetText(results64[0])
		case results65 := <-share.Ch65:
			searchTabItem65.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["65"], results65[1])
			searchTabWidget65.SetText(results65[0])
		case results66 := <-share.Ch66:
			searchTabItem66.Text = fmt.Sprintf("%v [%v]", parser.StandardAbbreviation["66"], results66[1])
			searchTabWidget66.SetText(results66[0])
		case results67 := <-share.Ch67:
			searchTabItem67.Text = fmt.Sprintf("... [%v]", results67[1])
			searchTabWidget67.SetText(results67[0])
		}
		progress += 100.0 / 67.0 / 100
		searchProgress.SetValue(progress)
	}
	searchProgress.SetValue(1.0)
	searchProgress.Hide()
	searchDisplayArea.Refresh()
}
