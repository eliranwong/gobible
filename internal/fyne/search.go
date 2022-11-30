package fyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/eliranwong/gobible/internal/parser"
	"github.com/eliranwong/gobible/internal/share"
)

var SearchTab1, SearchTab2, SearchTab3, SearchTab4, SearchTab5, SearchTab6, SearchTab7, SearchTab8, SearchTab9 *widget.Entry
var SearchTab10, SearchTab11, SearchTab12, SearchTab13, SearchTab14, SearchTab15, SearchTab16, SearchTab17, SearchTab18, SearchTab19 *widget.Entry
var SearchTab20, SearchTab21, SearchTab22, SearchTab23, SearchTab24, SearchTab25, SearchTab26, SearchTab27, SearchTab28, SearchTab29 *widget.Entry
var SearchTab30, SearchTab31, SearchTab32, SearchTab33, SearchTab34, SearchTab35, SearchTab36, SearchTab37, SearchTab38, SearchTab39 *widget.Entry
var SearchTab40, SearchTab41, SearchTab42, SearchTab43, SearchTab44, SearchTab45, SearchTab46, SearchTab47, SearchTab48, SearchTab49 *widget.Entry
var SearchTab50, SearchTab51, SearchTab52, SearchTab53, SearchTab54, SearchTab55, SearchTab56, SearchTab57, SearchTab58, SearchTab59 *widget.Entry
var SearchTab60, SearchTab61, SearchTab62, SearchTab63, SearchTab64, SearchTab65, SearchTab66, SearchTab67 *widget.Entry

func makeSearchTabs() fyne.CanvasObject {

	SearchTab1 = makeMultiLineEntry()
	SearchTab2 = makeMultiLineEntry()
	SearchTab3 = makeMultiLineEntry()
	SearchTab4 = makeMultiLineEntry()
	SearchTab5 = makeMultiLineEntry()
	SearchTab6 = makeMultiLineEntry()
	SearchTab7 = makeMultiLineEntry()
	SearchTab8 = makeMultiLineEntry()
	SearchTab9 = makeMultiLineEntry()
	SearchTab10 = makeMultiLineEntry()
	SearchTab11 = makeMultiLineEntry()
	SearchTab12 = makeMultiLineEntry()
	SearchTab13 = makeMultiLineEntry()
	SearchTab14 = makeMultiLineEntry()
	SearchTab15 = makeMultiLineEntry()
	SearchTab16 = makeMultiLineEntry()
	SearchTab17 = makeMultiLineEntry()
	SearchTab18 = makeMultiLineEntry()
	SearchTab19 = makeMultiLineEntry()
	SearchTab20 = makeMultiLineEntry()
	SearchTab21 = makeMultiLineEntry()
	SearchTab22 = makeMultiLineEntry()
	SearchTab23 = makeMultiLineEntry()
	SearchTab24 = makeMultiLineEntry()
	SearchTab25 = makeMultiLineEntry()
	SearchTab26 = makeMultiLineEntry()
	SearchTab27 = makeMultiLineEntry()
	SearchTab28 = makeMultiLineEntry()
	SearchTab29 = makeMultiLineEntry()
	SearchTab30 = makeMultiLineEntry()
	SearchTab31 = makeMultiLineEntry()
	SearchTab32 = makeMultiLineEntry()
	SearchTab33 = makeMultiLineEntry()
	SearchTab34 = makeMultiLineEntry()
	SearchTab35 = makeMultiLineEntry()
	SearchTab36 = makeMultiLineEntry()
	SearchTab37 = makeMultiLineEntry()
	SearchTab38 = makeMultiLineEntry()
	SearchTab39 = makeMultiLineEntry()
	SearchTab40 = makeMultiLineEntry()
	SearchTab41 = makeMultiLineEntry()
	SearchTab42 = makeMultiLineEntry()
	SearchTab43 = makeMultiLineEntry()
	SearchTab44 = makeMultiLineEntry()
	SearchTab45 = makeMultiLineEntry()
	SearchTab46 = makeMultiLineEntry()
	SearchTab47 = makeMultiLineEntry()
	SearchTab48 = makeMultiLineEntry()
	SearchTab49 = makeMultiLineEntry()
	SearchTab50 = makeMultiLineEntry()
	SearchTab51 = makeMultiLineEntry()
	SearchTab52 = makeMultiLineEntry()
	SearchTab53 = makeMultiLineEntry()
	SearchTab54 = makeMultiLineEntry()
	SearchTab55 = makeMultiLineEntry()
	SearchTab56 = makeMultiLineEntry()
	SearchTab57 = makeMultiLineEntry()
	SearchTab58 = makeMultiLineEntry()
	SearchTab59 = makeMultiLineEntry()
	SearchTab60 = makeMultiLineEntry()
	SearchTab61 = makeMultiLineEntry()
	SearchTab62 = makeMultiLineEntry()
	SearchTab63 = makeMultiLineEntry()
	SearchTab64 = makeMultiLineEntry()
	SearchTab65 = makeMultiLineEntry()
	SearchTab66 = makeMultiLineEntry()
	SearchTab67 = makeMultiLineEntry()

	SearchTabs := container.NewAppTabs(
		container.NewTabItem(parser.StandardAbbreviation["1"], SearchTab1),
		container.NewTabItem(parser.StandardAbbreviation["2"], SearchTab2),
		container.NewTabItem(parser.StandardAbbreviation["3"], SearchTab3),
		container.NewTabItem(parser.StandardAbbreviation["4"], SearchTab4),
		container.NewTabItem(parser.StandardAbbreviation["5"], SearchTab5),
		container.NewTabItem(parser.StandardAbbreviation["6"], SearchTab6),
		container.NewTabItem(parser.StandardAbbreviation["7"], SearchTab7),
		container.NewTabItem(parser.StandardAbbreviation["8"], SearchTab8),
		container.NewTabItem(parser.StandardAbbreviation["9"], SearchTab9),
		container.NewTabItem(parser.StandardAbbreviation["10"], SearchTab10),
		container.NewTabItem(parser.StandardAbbreviation["11"], SearchTab11),
		container.NewTabItem(parser.StandardAbbreviation["12"], SearchTab12),
		container.NewTabItem(parser.StandardAbbreviation["13"], SearchTab13),
		container.NewTabItem(parser.StandardAbbreviation["14"], SearchTab14),
		container.NewTabItem(parser.StandardAbbreviation["15"], SearchTab15),
		container.NewTabItem(parser.StandardAbbreviation["16"], SearchTab16),
		container.NewTabItem(parser.StandardAbbreviation["17"], SearchTab17),
		container.NewTabItem(parser.StandardAbbreviation["18"], SearchTab18),
		container.NewTabItem(parser.StandardAbbreviation["19"], SearchTab19),
		container.NewTabItem(parser.StandardAbbreviation["20"], SearchTab20),
		container.NewTabItem(parser.StandardAbbreviation["21"], SearchTab21),
		container.NewTabItem(parser.StandardAbbreviation["22"], SearchTab22),
		container.NewTabItem(parser.StandardAbbreviation["23"], SearchTab23),
		container.NewTabItem(parser.StandardAbbreviation["24"], SearchTab24),
		container.NewTabItem(parser.StandardAbbreviation["25"], SearchTab25),
		container.NewTabItem(parser.StandardAbbreviation["26"], SearchTab26),
		container.NewTabItem(parser.StandardAbbreviation["27"], SearchTab27),
		container.NewTabItem(parser.StandardAbbreviation["28"], SearchTab28),
		container.NewTabItem(parser.StandardAbbreviation["29"], SearchTab29),
		container.NewTabItem(parser.StandardAbbreviation["30"], SearchTab30),
		container.NewTabItem(parser.StandardAbbreviation["31"], SearchTab31),
		container.NewTabItem(parser.StandardAbbreviation["32"], SearchTab32),
		container.NewTabItem(parser.StandardAbbreviation["33"], SearchTab33),
		container.NewTabItem(parser.StandardAbbreviation["34"], SearchTab34),
		container.NewTabItem(parser.StandardAbbreviation["35"], SearchTab35),
		container.NewTabItem(parser.StandardAbbreviation["36"], SearchTab36),
		container.NewTabItem(parser.StandardAbbreviation["37"], SearchTab37),
		container.NewTabItem(parser.StandardAbbreviation["38"], SearchTab38),
		container.NewTabItem(parser.StandardAbbreviation["39"], SearchTab39),
		container.NewTabItem(parser.StandardAbbreviation["40"], SearchTab40),
		container.NewTabItem(parser.StandardAbbreviation["41"], SearchTab41),
		container.NewTabItem(parser.StandardAbbreviation["42"], SearchTab42),
		container.NewTabItem(parser.StandardAbbreviation["43"], SearchTab43),
		container.NewTabItem(parser.StandardAbbreviation["44"], SearchTab44),
		container.NewTabItem(parser.StandardAbbreviation["45"], SearchTab45),
		container.NewTabItem(parser.StandardAbbreviation["46"], SearchTab46),
		container.NewTabItem(parser.StandardAbbreviation["47"], SearchTab47),
		container.NewTabItem(parser.StandardAbbreviation["48"], SearchTab48),
		container.NewTabItem(parser.StandardAbbreviation["49"], SearchTab49),
		container.NewTabItem(parser.StandardAbbreviation["50"], SearchTab50),
		container.NewTabItem(parser.StandardAbbreviation["51"], SearchTab51),
		container.NewTabItem(parser.StandardAbbreviation["52"], SearchTab52),
		container.NewTabItem(parser.StandardAbbreviation["53"], SearchTab53),
		container.NewTabItem(parser.StandardAbbreviation["54"], SearchTab54),
		container.NewTabItem(parser.StandardAbbreviation["55"], SearchTab55),
		container.NewTabItem(parser.StandardAbbreviation["56"], SearchTab56),
		container.NewTabItem(parser.StandardAbbreviation["57"], SearchTab57),
		container.NewTabItem(parser.StandardAbbreviation["58"], SearchTab58),
		container.NewTabItem(parser.StandardAbbreviation["59"], SearchTab59),
		container.NewTabItem(parser.StandardAbbreviation["60"], SearchTab60),
		container.NewTabItem(parser.StandardAbbreviation["61"], SearchTab61),
		container.NewTabItem(parser.StandardAbbreviation["62"], SearchTab62),
		container.NewTabItem(parser.StandardAbbreviation["63"], SearchTab63),
		container.NewTabItem(parser.StandardAbbreviation["64"], SearchTab64),
		container.NewTabItem(parser.StandardAbbreviation["65"], SearchTab65),
		container.NewTabItem(parser.StandardAbbreviation["66"], SearchTab66),
		container.NewTabItem("...", makeMultiLineEntry()),
	)
	SearchTabs.SetTabLocation(container.TabLocationTop)
	return container.NewBorder(nil, nil, nil, nil, SearchTabs)
}

func populateSearchTabs() {
	// populate tabs as soon as available
	for i := 1; i <= 67; i++ {
		select {
		case results1 := <-share.Ch1:
			SearchTab1.SetText(results1)
		case results2 := <-share.Ch2:
			SearchTab2.SetText(results2)
		case results3 := <-share.Ch3:
			SearchTab3.SetText(results3)
		case results4 := <-share.Ch4:
			SearchTab4.SetText(results4)
		case results5 := <-share.Ch5:
			SearchTab5.SetText(results5)
		case results6 := <-share.Ch6:
			SearchTab6.SetText(results6)
		case results7 := <-share.Ch7:
			SearchTab7.SetText(results7)
		case results8 := <-share.Ch8:
			SearchTab8.SetText(results8)
		case results9 := <-share.Ch9:
			SearchTab9.SetText(results9)
		case results10 := <-share.Ch10:
			SearchTab10.SetText(results10)
		case results11 := <-share.Ch11:
			SearchTab11.SetText(results11)
		case results12 := <-share.Ch12:
			SearchTab12.SetText(results12)
		case results13 := <-share.Ch13:
			SearchTab13.SetText(results13)
		case results14 := <-share.Ch14:
			SearchTab14.SetText(results14)
		case results15 := <-share.Ch15:
			SearchTab15.SetText(results15)
		case results16 := <-share.Ch16:
			SearchTab16.SetText(results16)
		case results17 := <-share.Ch17:
			SearchTab17.SetText(results17)
		case results18 := <-share.Ch18:
			SearchTab18.SetText(results18)
		case results19 := <-share.Ch19:
			SearchTab19.SetText(results19)
		case results20 := <-share.Ch20:
			SearchTab20.SetText(results20)
		case results21 := <-share.Ch21:
			SearchTab21.SetText(results21)
		case results22 := <-share.Ch22:
			SearchTab22.SetText(results22)
		case results23 := <-share.Ch23:
			SearchTab23.SetText(results23)
		case results24 := <-share.Ch24:
			SearchTab24.SetText(results24)
		case results25 := <-share.Ch25:
			SearchTab25.SetText(results25)
		case results26 := <-share.Ch26:
			SearchTab26.SetText(results26)
		case results27 := <-share.Ch27:
			SearchTab27.SetText(results27)
		case results28 := <-share.Ch28:
			SearchTab28.SetText(results28)
		case results29 := <-share.Ch29:
			SearchTab29.SetText(results29)
		case results30 := <-share.Ch30:
			SearchTab30.SetText(results30)
		case results31 := <-share.Ch31:
			SearchTab31.SetText(results31)
		case results32 := <-share.Ch32:
			SearchTab32.SetText(results32)
		case results33 := <-share.Ch33:
			SearchTab33.SetText(results33)
		case results34 := <-share.Ch34:
			SearchTab34.SetText(results34)
		case results35 := <-share.Ch35:
			SearchTab35.SetText(results35)
		case results36 := <-share.Ch36:
			SearchTab36.SetText(results36)
		case results37 := <-share.Ch37:
			SearchTab37.SetText(results37)
		case results38 := <-share.Ch38:
			SearchTab38.SetText(results38)
		case results39 := <-share.Ch39:
			SearchTab39.SetText(results39)
		case results40 := <-share.Ch40:
			SearchTab40.SetText(results40)
		case results41 := <-share.Ch41:
			SearchTab41.SetText(results41)
		case results42 := <-share.Ch42:
			SearchTab42.SetText(results42)
		case results43 := <-share.Ch43:
			SearchTab43.SetText(results43)
		case results44 := <-share.Ch44:
			SearchTab44.SetText(results44)
		case results45 := <-share.Ch45:
			SearchTab45.SetText(results45)
		case results46 := <-share.Ch46:
			SearchTab46.SetText(results46)
		case results47 := <-share.Ch47:
			SearchTab47.SetText(results47)
		case results48 := <-share.Ch48:
			SearchTab48.SetText(results48)
		case results49 := <-share.Ch49:
			SearchTab49.SetText(results49)
		case results50 := <-share.Ch50:
			SearchTab50.SetText(results50)
		case results51 := <-share.Ch51:
			SearchTab51.SetText(results51)
		case results52 := <-share.Ch52:
			SearchTab52.SetText(results52)
		case results53 := <-share.Ch53:
			SearchTab53.SetText(results53)
		case results54 := <-share.Ch54:
			SearchTab54.SetText(results54)
		case results55 := <-share.Ch55:
			SearchTab55.SetText(results55)
		case results56 := <-share.Ch56:
			SearchTab56.SetText(results56)
		case results57 := <-share.Ch57:
			SearchTab57.SetText(results57)
		case results58 := <-share.Ch58:
			SearchTab58.SetText(results58)
		case results59 := <-share.Ch59:
			SearchTab59.SetText(results59)
		case results60 := <-share.Ch60:
			SearchTab60.SetText(results60)
		case results61 := <-share.Ch61:
			SearchTab61.SetText(results61)
		case results62 := <-share.Ch62:
			SearchTab62.SetText(results62)
		case results63 := <-share.Ch63:
			SearchTab63.SetText(results63)
		case results64 := <-share.Ch64:
			SearchTab64.SetText(results64)
		case results65 := <-share.Ch65:
			SearchTab65.SetText(results65)
		case results66 := <-share.Ch66:
			SearchTab66.SetText(results66)
		case results67 := <-share.Ch67:
			SearchTab67.SetText(results67)
		}
	}
}
