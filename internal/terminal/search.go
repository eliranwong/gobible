package terminal

import (
	"fmt"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/eliranwong/gobible/internal/parser"
	"github.com/eliranwong/gobible/internal/share"
	"github.com/schollz/progressbar/v3"
)

func showSearchResults() {
	searchProgress := progressbar.Default(10000)
	step := 10000.0 / 67.0
	progressIncrement := int(step)
	start := time.Now()

	var results1, results2, results3, results4, results5, results6, results7, results8, results9 [][]string
	var results10, results11, results12, results13, results14, results15, results16, results17, results18, results19 [][]string
	var results20, results21, results22, results23, results24, results25, results26, results27, results28, results29 [][]string
	var results30, results31, results32, results33, results34, results35, results36, results37, results38, results39 [][]string
	var results40, results41, results42, results43, results44, results45, results46, results47, results48, results49 [][]string
	var results50, results51, results52, results53, results54, results55, results56, results57, results58, results59 [][]string
	var results60, results61, results62, results63, results64, results65, results66, results67 [][]string

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
		searchProgress.Add(progressIncrement)
	}
	elapsed := time.Since(start)
	// finish progress bar
	searchProgress.Add(17)

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

	labels := []string{
		"~ exit ~",
	}
	labels = append(labels, resultLabelAll)
	labels = append(labels, resultLabel1)
	labels = append(labels, resultLabel2)
	labels = append(labels, resultLabel3)
	labels = append(labels, resultLabel4)
	labels = append(labels, resultLabel5)
	labels = append(labels, resultLabel6)
	labels = append(labels, resultLabel7)
	labels = append(labels, resultLabel8)
	labels = append(labels, resultLabel9)
	labels = append(labels, resultLabel10)
	labels = append(labels, resultLabel11)
	labels = append(labels, resultLabel12)
	labels = append(labels, resultLabel13)
	labels = append(labels, resultLabel14)
	labels = append(labels, resultLabel15)
	labels = append(labels, resultLabel16)
	labels = append(labels, resultLabel17)
	labels = append(labels, resultLabel18)
	labels = append(labels, resultLabel19)
	labels = append(labels, resultLabel20)
	labels = append(labels, resultLabel21)
	labels = append(labels, resultLabel22)
	labels = append(labels, resultLabel23)
	labels = append(labels, resultLabel24)
	labels = append(labels, resultLabel25)
	labels = append(labels, resultLabel26)
	labels = append(labels, resultLabel27)
	labels = append(labels, resultLabel28)
	labels = append(labels, resultLabel29)
	labels = append(labels, resultLabel30)
	labels = append(labels, resultLabel31)
	labels = append(labels, resultLabel32)
	labels = append(labels, resultLabel33)
	labels = append(labels, resultLabel34)
	labels = append(labels, resultLabel35)
	labels = append(labels, resultLabel36)
	labels = append(labels, resultLabel37)
	labels = append(labels, resultLabel38)
	labels = append(labels, resultLabel39)
	labels = append(labels, resultLabel40)
	labels = append(labels, resultLabel41)
	labels = append(labels, resultLabel42)
	labels = append(labels, resultLabel43)
	labels = append(labels, resultLabel44)
	labels = append(labels, resultLabel45)
	labels = append(labels, resultLabel46)
	labels = append(labels, resultLabel47)
	labels = append(labels, resultLabel48)
	labels = append(labels, resultLabel49)
	labels = append(labels, resultLabel50)
	labels = append(labels, resultLabel51)
	labels = append(labels, resultLabel52)
	labels = append(labels, resultLabel53)
	labels = append(labels, resultLabel54)
	labels = append(labels, resultLabel55)
	labels = append(labels, resultLabel56)
	labels = append(labels, resultLabel57)
	labels = append(labels, resultLabel58)
	labels = append(labels, resultLabel59)
	labels = append(labels, resultLabel60)
	labels = append(labels, resultLabel61)
	labels = append(labels, resultLabel62)
	labels = append(labels, resultLabel63)
	labels = append(labels, resultLabel64)
	labels = append(labels, resultLabel65)
	labels = append(labels, resultLabel66)
	labels = append(labels, resultLabel67)

	promptResult := func() string {
		var qs = []*survey.Question{
			{
				Name: "Book",
				Prompt: &survey.Select{
					Message: "Choose to display search results:",
					Options: labels,
					Default: resultLabelAll,
					Help:    "Choose '~ exit ~' to exit",
				},
			},
		}
		answers := struct {
			Book string
		}{}
		// perform the questions
		err := survey.Ask(qs, &answers, pageSize)
		if err != nil {
			fmt.Println(err.Error())
			return ""
		}
		if err != nil {
			fmt.Println(err.Error())
			return ""
		}
		return answers.Book
	}

	for {
		book := promptResult()
		if book == "~ exit ~" || book == "" {
			break
		} else {
			switch book {
			case resultLabelAll:
				displayResults(allResults)
			case resultLabel1:
				displayResults(results1)
			case resultLabel2:
				displayResults(results2)
			case resultLabel3:
				displayResults(results3)
			case resultLabel4:
				displayResults(results4)
			case resultLabel5:
				displayResults(results5)
			case resultLabel6:
				displayResults(results6)
			case resultLabel7:
				displayResults(results7)
			case resultLabel8:
				displayResults(results8)
			case resultLabel9:
				displayResults(results9)
			case resultLabel10:
				displayResults(results10)
			case resultLabel11:
				displayResults(results11)
			case resultLabel12:
				displayResults(results12)
			case resultLabel13:
				displayResults(results13)
			case resultLabel14:
				displayResults(results14)
			case resultLabel15:
				displayResults(results15)
			case resultLabel16:
				displayResults(results16)
			case resultLabel17:
				displayResults(results17)
			case resultLabel18:
				displayResults(results18)
			case resultLabel19:
				displayResults(results19)
			case resultLabel20:
				displayResults(results20)
			case resultLabel21:
				displayResults(results21)
			case resultLabel22:
				displayResults(results22)
			case resultLabel23:
				displayResults(results23)
			case resultLabel24:
				displayResults(results24)
			case resultLabel25:
				displayResults(results25)
			case resultLabel26:
				displayResults(results26)
			case resultLabel27:
				displayResults(results27)
			case resultLabel28:
				displayResults(results28)
			case resultLabel29:
				displayResults(results29)
			case resultLabel30:
				displayResults(results30)
			case resultLabel31:
				displayResults(results31)
			case resultLabel32:
				displayResults(results32)
			case resultLabel33:
				displayResults(results33)
			case resultLabel34:
				displayResults(results34)
			case resultLabel35:
				displayResults(results35)
			case resultLabel36:
				displayResults(results36)
			case resultLabel37:
				displayResults(results37)
			case resultLabel38:
				displayResults(results38)
			case resultLabel39:
				displayResults(results39)
			case resultLabel40:
				displayResults(results40)
			case resultLabel41:
				displayResults(results41)
			case resultLabel42:
				displayResults(results42)
			case resultLabel43:
				displayResults(results43)
			case resultLabel44:
				displayResults(results44)
			case resultLabel45:
				displayResults(results45)
			case resultLabel46:
				displayResults(results46)
			case resultLabel47:
				displayResults(results47)
			case resultLabel48:
				displayResults(results48)
			case resultLabel49:
				displayResults(results49)
			case resultLabel50:
				displayResults(results50)
			case resultLabel51:
				displayResults(results51)
			case resultLabel52:
				displayResults(results52)
			case resultLabel53:
				displayResults(results53)
			case resultLabel54:
				displayResults(results54)
			case resultLabel55:
				displayResults(results55)
			case resultLabel56:
				displayResults(results56)
			case resultLabel57:
				displayResults(results57)
			case resultLabel58:
				displayResults(results58)
			case resultLabel59:
				displayResults(results59)
			case resultLabel60:
				displayResults(results60)
			case resultLabel61:
				displayResults(results61)
			case resultLabel62:
				displayResults(results62)
			case resultLabel63:
				displayResults(results63)
			case resultLabel64:
				displayResults(results64)
			case resultLabel65:
				displayResults(results65)
			case resultLabel66:
				displayResults(results66)
			case resultLabel67:
				displayResults(results67)
			}
		}
	}
	// report search time
	fmt.Printf("search took %s\n", elapsed)
}

func displayResults(results [][]string) {
	var display strings.Builder
	for _, element := range results {
		display.WriteString(fmt.Sprintf("(%v) %v\n", element[1], element[2]))
	}
	displayOnTerminal(display.String())
}
