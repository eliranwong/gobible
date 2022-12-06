package terminal

import (
	"fmt"
	"path/filepath"
	"sort"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/eliranwong/gobible/internal/parser"
	"github.com/eliranwong/gobible/internal/share"
	"github.com/eliranwong/gobible/internal/shortcuts"
)

func getAbbreviations(toComplete string) []string {
	var books []int
	var abbreviations []string
	for key, _ := range parser.StandardAbbreviation {
		b, _ := strconv.Atoi(key)
		books = append(books, b)
	}
	sort.Ints(books)
	for _, b := range books {
		abbreviations = append(abbreviations, parser.StandardAbbreviation[strconv.Itoa(b)])
	}
	return abbreviations
}

func promptBible() {
	bibles, _ := shortcuts.WalkMatch(filepath.Join(share.Data, filepath.FromSlash("bibles")), "*.bible", true)
	var qs = []*survey.Question{
		{
			Name: "Module",
			Prompt: &survey.Select{
				Message: "Choose a version:",
				Options: bibles,
				Default: share.Bible,
			},
		},
		{
			Name: "Reference",
			Prompt: &survey.Input{
				Message: "Enter a reference:",
				Suggest: getAbbreviations,
			},
			Validate: survey.Required,
		},
	}
	answers := struct {
		Module    string
		Reference string
	}{}
	// perform the questions
	err := survey.Ask(qs, &answers, pageSize)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if isValidEntry(answers.Reference) {
		RunCommand(answers.Reference, answers.Module)
	}
}
