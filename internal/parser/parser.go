// Package parser
// parse bible reference
package parser

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/eliranwong/gobible/internal/regex"
	"github.com/eliranwong/gobible/internal/shortcuts"
)

// convert b c v integers to verse reference string
func BcvToVerseReference(bcvList []int) string {
	b := bcvList[0]
	c := bcvList[1]
	v := bcvList[2]
	var c2, v2 int
	if len(bcvList) == 5 {
		c2 = bcvList[3]
		v2 = bcvList[4]
	}
	reference := "BOOK 0:0"
	bookNoStr := strconv.Itoa(b)
	abbreviation, knownBookNo := StandardAbbreviation[bookNoStr]
	if knownBookNo {
		if (c2 != 0) && (c2 == c) && (v2 > v) {
			reference = fmt.Sprintf("%v %v:%v-%v", abbreviation, c, v, v2)
		} else if (c2 != 0) && (c2 > c) {
			reference = fmt.Sprintf("%v %v:%v-%v:%v", abbreviation, c, v, c2, v2)
		} else {
			reference = fmt.Sprintf("%v %v:%v", abbreviation, c, v)
		}
	}
	return reference
}

func ParseText(text string) string {
	// add a space at the end of the text, to avoid indefinite loop in later steps
	// this extra space will be removed when parsing is finished.
	taggedText := text + " "

	var p string
	var searchReplace [][2]string

	// remove bcv tags, if any, to avoid duplication of tagging in later steps
	p = `<ref onclick="bcv\([^\(\)]*?\)">(.*?)</ref>`
	searchReplace = [][2]string{
		{p, "$1"},
	}
	regex.ReplaceAllStringLoop(text, "m", p, searchReplace)

	// sorting books by length
	sortedBooks := shortcuts.MapKeysToStringSlice(BibleBookNo)
	sort.Slice(sortedBooks, func(i, j int) bool {
		return len(sortedBooks[i]) > len(sortedBooks[j])
	})

	var bookString, bookNumber string
	for _, book := range sortedBooks {
		bookString = book

		searchReplace = [][2]string{
			{`\.`, `[\.]*?`},           // make dot "." optional for an abbreviation
			{`^([0-9]+?) `, `$1[ ]*?`}, // make space " " optional in some cases
			{`^([I]+?) `, `$1[ ]*?`},
			{`^(IV) `, `$1[ ]*?`},
		}
		bookString = regex.ReplaceAllString(bookString, "", searchReplace)

		// get assigned book number from dictionary
		bookNumber = BibleBookNo[book]

		// search & replace for marking book
		searchReplace = [][2]string{
			{fmt.Sprintf(`(%v) ([0-9])`, bookString), fmt.Sprintf(`『%v｜$1』 $2`, bookNumber)},
		}
		taggedText = regex.ReplaceAllString(taggedText, "m", searchReplace)
	}
	// check
	//fmt.Println(1, taggedText)

	searchReplace = [][2]string{
		// add first set of taggings:
		{
			`『([0-9]+?)｜([^『』]*?)』 ([0-9]+?):([0-9]+?)([^0-9])`,
			`<ref onclick="bcv($1,$3,$4)">$2 $3:$4</ref｝$5`,
		},
		{
			`『([0-9]+?)｜([^『』]*?)』 ([0-9]+?)([^0-9])`,
			`<ref onclick="bcv($1,$3,)">$2 $3</ref｝$4`,
		},
		// fix references without verse numbers
		// fix books with chapter 1 ONLY; oneChapterBook = [31,57,63,64,65,72,73,75,79,85]
		{
			`<ref onclick="bcv\((31|57|63|64|65|72|73|75|79|85),([0-9]+?),\)">`,
			`<ref onclick="bcv($1,1,$2)">`,
		},
		// fix references of chapters without verse number; assign verse number 1 in taggings
		{
			`<ref onclick="bcv\(([0-9]+?),([0-9]+?),\)">`,
			`<ref onclick="bcv($1,$2,1)">＊`,
		},
	}
	taggedText = regex.ReplaceAllString(taggedText, "m", searchReplace)
	// check
	//fmt.Println(2, taggedText)

	// check if verses following tagged references, e.g. Book 1:1-2:1; 3:2-4, 5; Jude 1
	searchReplace = [][2]string{
		{
			`<ref onclick="bcv\(([0-9]+?),([0-9]+?),([0-9]+?)\)">([^｝]*?)</ref｝([,-–;][ ]*?)([0-9]+?):([0-9]+?)([^0-9])`,
			`<ref onclick="bcv($1,$2,$3)">$4</ref｝$5<ref onclick="bcv($1,$6,$7)">$6:$7</ref｝$8`,
		},
		{
			`<ref onclick="bcv\(([0-9]+?),([0-9]+?),([0-9]+?)\)">([^＊][^｝]*?)</ref｝([,-–;][ ]*?)([0-9]+?)([^:0-9])`,
			`<ref onclick="bcv($1,$2,$3)">$4</ref｝$5<ref onclick="bcv($1,$2,$6)">$6</ref｝$7`,
		},
		{
			`<ref onclick="bcv\(([0-9]+?),([0-9]+?),([0-9]+?)\)">([^＊][^｝]*?)</ref｝([,-–;][ ]*?)([0-9]+?):([^0-9])`,
			`<ref onclick="bcv($1,$2,$3)">$4</ref｝$5<ref onclick="bcv($1,$2,$6)">$6</ref｝:$7`,
		},
		{
			`<ref onclick="bcv\(([0-9]+?),([0-9]+?),([0-9]+?)\)">(＊[^｝]*?)</ref｝([,-–;][ ]*?)([0-9]+?)([^:0-9])`,
			`<ref onclick="bcv($1,$2,$3)">$4</ref｝$5<ref onclick="bcv($1,$6,1)">＊$6</ref｝$7`,
		},
		{
			`<ref onclick="bcv\(([0-9]+?),([0-9]+?),([0-9]+?)\)">(＊[^｝]*?)</ref｝([,-–;][ ]*?)([0-9]+?):([^0-9])`,
			`<ref onclick="bcv($1,$2,$3)">$4</ref｝$5<ref onclick="bcv($1,$6,1)">＊$6</ref｝:$7`,
		},
	}
	taggedText = regex.ReplaceAllStringLoop(taggedText, "m", `</ref｝[,-–;][ ]*?[0-9]`, searchReplace)
	// check
	//fmt.Println(3, taggedText)

	// clear special markers
	searchReplace = [][2]string{
		{`『[0-9]+?|([^『』]*?)』`, `$1`},
		{`(<ref onclick="bcv\([0-9]+?,[0-9]+?,[0-9]+?\)">)＊`, `$1`},
		{`</ref｝`, `</ref>`},
	}
	taggedText = regex.ReplaceAllString(taggedText, "m", searchReplace)
	// check
	//fmt.Println(4, taggedText)

	// handling range of verses
	// e.g. John 3:16 is tagged as <ref onclick="bcv(43,3,16)">John 3:16</ref>
	// e.g. John 3:14-16 is tagged as <ref onclick="bcv(43,3,14,3,16)">John 3:14-16</ref>
	// e.g. John 3:14-4:3 is tagged as <ref onclick="bcv(43,3,14,4,3)">John 3:14-4:3</ref>

	loopPattern := `<ref onclick="bcv\(([0-9]+?),([0-9]+?),([0-9]+?)\)">([^<>]*?)</ref>([-–])<ref onclick="bcv\(([0-9]+?),([0-9]+?),([0-9]+?)\)">`
	compiledPattern := regexp.MustCompile(fmt.Sprintf(`(?%v)%v`, "m", loopPattern))
	for compiledPattern.MatchString(taggedText) {
		taggedText = compiledPattern.ReplaceAllStringFunc(taggedText, func(text string) string {
			text = strings.ReplaceAll(text, "–", "-")
			references := strings.Split(text, "ref>-<ref")
			reference0 := references[0] + "ref>"
			reference1 := "<ref" + references[1]
			p := regexp.MustCompile(`^(<ref onclick="bcv\([0-9]+?,).*?$`)
			if p.ReplaceAllString(reference0, "$1") == p.ReplaceAllString(reference1, "$1") {
				pp := regexp.MustCompile(`<ref onclick="bcv\(([0-9]+?),([0-9]+?),([0-9]+?)\)">([^<>]*?)</ref>([-–])<ref onclick="bcv\([0-9]+?,([0-9]+?),([0-9]+?)\)">`)
				text = pp.ReplaceAllString(text, `<ref onclick="bcv($1,$2,$3,$6,$7)">$4$5`)
			} else {
				// avoid infinite loop
				text = fmt.Sprintf("%v--%v", reference0, reference1)
			}
			return text
		})
	}

	// remove the extra space, added at the beginning of this function
	taggedText = strings.TrimSpace(taggedText)

	return taggedText
}

func ExtractAllReferences(text string, tagged bool) [][]int {
	// Parse the text only if it is not tagged.
	if !tagged {
		text = ParseText(text)
	}
	verseReferenceList := [][]int{}
	searchPattern := regexp.MustCompile(`bcv\(([0-9]+?,[0-9]+?,[0-9]+?[^\)\(]*?)\)`)
	for _, v := range searchPattern.FindAllStringSubmatch(text, -1) {
		bcv := []int{}
		for _, v2 := range strings.Split(v[1], ",") {
			number, _ := strconv.Atoi(v2)
			bcv = append(bcv, number)
		}
		verseReferenceList = append(verseReferenceList, bcv)
	}
	return verseReferenceList
}

func BookNameToNumber(name string) int {
	numberStr, ok := BibleBookNo[name]
	if ok {
		number, _ := strconv.Atoi(numberStr)
		return number
	} else {
		numberStr, ok := BibleBookNo[(name + ".")]
		if ok {
			number, _ := strconv.Atoi(numberStr)
			return number
		}
	}
	return 0

}

func BookNumberToName(number int) string {
	numberStr := strconv.Itoa(number)
	name, ok := StandardBookname[numberStr]
	if ok {
		return name
	}
	return ""
}

func BookNumberToAbb(number int) string {
	numberStr := strconv.Itoa(number)
	abb, ok := StandardAbbreviation[numberStr]
	if ok {
		return abb
	}
	return ""
}
