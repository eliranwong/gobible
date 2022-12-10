package terminal

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/eliranwong/gobible/internal/bible"
	"github.com/eliranwong/gobible/internal/parser"
	"github.com/eliranwong/gobible/internal/share"
)

func promptChapter(bookName string) {
	bookNo := parser.BookNameToNumber(bookName)
	chapters := bible.GetChapters(share.Bible, bookNo)
	chapterStrings := share.IntSliceToStringSlice(chapters)
	var qs = []*survey.Question{
		{
			Name: "Chapter",
			Prompt: &survey.Select{
				Message: "Choose a chapter:",
				Options: chapterStrings,
			},
		},
	}
	answers := struct {
		Chapter string
	}{}
	// perform the questions
	err := survey.Ask(qs, &answers, pageSize)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	RunCommand(fmt.Sprintf("%v %v", bookName, answers.Chapter), share.Bible)
}
