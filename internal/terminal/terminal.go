// terminal mode interface
package terminal

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/common-nighthawk/go-figure"
	"github.com/eliranwong/gobible/internal/bible"
	"github.com/eliranwong/gobible/internal/parser"
	"github.com/eliranwong/gobible/internal/share"
)

var cancel string = ".cancel"
var pageSize survey.AskOpt = survey.WithPageSize(15)

var commands map[string]func() = map[string]func(){
	//".cancel": func() { isValidEntry(cancel) },
	".bible":  promptBible,
	".search": promptSearch,
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") //Windows example, its tested
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func getCommands(toComplete string) []string {
	var commandKeys []string
	for key := range commands {
		commandKeys = append(commandKeys, key)
	}
	sort.Strings(commandKeys)
	return commandKeys
}

func header() {
	share.Divider()
	figure.NewFigure("GoBible", "", true).Print()
	message := fmt.Sprintf("\nCurrent reference: %v", share.Info(share.Reference))
	fmt.Println(message)
}

func Terminal() {
	clearScreen()
	share.SetupData()
	//check theme color
	//share.TestThemeColors()
	for {
		header()
		command := promptCommand()
		if command == ".quit" {
			break
		} else if val, ok := commands[command]; ok {
			val()
		} else {
			RunCommand(command, share.Bible)
		}
	}
}

func promptCommand() string {
	// the questions to ask
	var q = []*survey.Question{
		{
			Name: "Command",
			Prompt: &survey.Input{
				Message: ">>>",
				Suggest: getCommands,
				Help:    "enter a bible reference to open / a GoBible command to quit / '.quit' to quit [ctrl+c]",
			},
			Validate: survey.Required,
		},
	}
	answers := struct {
		Command string
	}{}

	// ask the question
	err := survey.Ask(q, &answers, pageSize)

	if err != nil {
		if err == terminal.InterruptErr {
			return ".quit"
		} else {
			fmt.Println(err.Error())
		}
		return ""
	}
	return answers.Command
}

func RunCommand(command, bibleModule string) {
	if !(strings.TrimSpace(command) == "") {
		references := parser.ExtractAllReferences(command, false)
		// search bible when there is no valid bible reference
		if len(references) == 0 {
			runSearch(share.SearchMethod, share.Bible, command)
		} else {
			// reset bible chapter for display
			bible.Chapter.Reset()
			bible.Read(bibleModule, references)
			displayOnTerminal(bible.Chapter.String())
		}
	}
}

func isValidEntry(entry string) bool {
	if strings.TrimSpace(entry) == "" {
		return false
	} else if entry == cancel {
		fmt.Println(share.Debug("Action cancelled!"))
		return false
	}
	return true
}

func displayOnTerminal(text string) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("more")
	} else {
		// add -R to support color text
		cmd = exec.Command("less", "-R")
	}
	// Feed it with the string you want to display.
	cmd.Stdin = strings.NewReader(text)
	// This is crucial - otherwise it will write to a null device.
	cmd.Stdout = os.Stdout
	// Fork off a process and wait for it to terminate.
	err := cmd.Run()
	if err != nil {
		//fmt.Println(err)
		//log.Fatal(err)
		// fallback to simply use print
		fmt.Println(text)
	}
}
