package terminal

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/eliranwong/gobible/internal/bible"
	"github.com/eliranwong/gobible/internal/share"
)

func promptVerses(message, selected string, options, descriptions []string, run func(string)) {
	var qs = []*survey.Question{
		{
			Name: "Verse",
			Prompt: &survey.Select{
				Message:     message,
				Options:     options,
				Description: func(value string, index int) string { return descriptions[index] },
				Help:        "Select to open",
				Default:     selected,
			},
		},
	}
	answers := struct {
		Verse string
	}{}
	// perform the questions
	err := survey.Ask(qs, &answers, pageSize)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	run(answers.Verse)
}

func crossreference() {
	run := func(s string) {
		l := strings.Split(s, " [")
		reference, module := l[0], l[1]
		RunCommand(reference, module[:len(module)-1])
	}
	message := fmt.Sprintf("Cross-reference: %v", share.Reference)
	options, descriptions := bible.GetXrefs2(share.Reference)
	promptVerses(message, options[0], options, descriptions, run)
}
