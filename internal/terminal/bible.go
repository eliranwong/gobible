package terminal

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/eliranwong/gobible/internal/bible"
	"github.com/eliranwong/gobible/internal/share"
)

func promptBible() {
	var qs = []*survey.Question{
		{
			Name: "Module",
			Prompt: &survey.Select{
				Message: "Choose a version:",
				Options: share.Bibles,
				Default: share.Bible,
			},
		},
		{
			Name: "Reference",
			Prompt: &survey.Input{
				Message: "Enter a reference:",
				Suggest: func(toComplete string) []string {
					return bible.GetBookNames(share.Bible)
				},
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
	if isValidEntry(answers.Reference) {
		RunCommand(answers.Reference, answers.Module)
	}
}
