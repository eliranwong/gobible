package share

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gookit/color"
)

func TestThemeColors() {
	// Info color style
	fmt.Println(Info("Info test"))
	// Note color style
	fmt.Println(Note("Note test"))
	// Warn color style
	fmt.Println(Warn("Warn test"))
	// Light color style
	fmt.Println(Light("Light test"))
	// Error color style
	fmt.Println(Error("Error test"))
	// Danger color style
	fmt.Println(Danger("Danger test"))
	// Debug color style
	fmt.Println(Debug("Debug test"))
	// Notice color style
	fmt.Println(Notice("Notice test"))
	// Comment color style
	fmt.Println(Comment("Comment test"))
	// Success color style
	fmt.Println(Success("Success test"))
	// Primary color style
	fmt.Println(Primary("Primary test"))
	// Question color style
	fmt.Println(Question("Question test"))
	// Secondary color style
	fmt.Println(Secondary("Secondary test"))
}

// Info color style
// Info = &Theme{"info", Style{OpReset, FgGreen}}
func Info(s string) string {
	if Mode == "" {
		return color.Info.Render(s)
	}
	return s
}

// Note color style
// Note = &Theme{"note", Style{OpBold, FgLightCyan}}
func Note(s string) string {
	if Mode == "" {
		return color.Note.Render(s)
	}
	return s
}

// Warn color style
// Warn = &Theme{"warning", Style{OpBold, FgYellow}}
func Warn(s string) string {
	if Mode == "" {
		return color.Warn.Render(s)
	}
	return s
}

// Light color style
// Light = &Theme{"light", Style{FgLightWhite, BgBlack}}
func Light(s string) string {
	if Mode == "" {
		return color.Light.Render(s)
	}
	return s
}

// Error color style
// Error = &Theme{"error", Style{FgLightWhite, BgRed}}
func Error(s string) string {
	if Mode == "" {
		return color.Error.Render(s)
	}
	return s
}

// Danger color style
// Danger = &Theme{"danger", Style{OpBold, FgRed}}
func Danger(s string) string {
	if Mode == "" {
		return color.Danger.Render(s)
	}
	return s
}

// Debug color style
// Debug = &Theme{"debug", Style{OpReset, FgCyan}}
func Debug(s string) string {
	if Mode == "" {
		return color.Debug.Render(s)
	}
	return s
}

// Notice color style
// Notice = &Theme{"notice", Style{OpBold, FgCyan}}
func Notice(s string) string {
	if Mode == "" {
		return color.Notice.Render(s)
	}
	return s
}

// Comment color style
// Comment = &Theme{"comment", Style{OpReset, FgYellow}}
func Comment(s string) string {
	if Mode == "" {
		return color.Comment.Render(s)
	}
	return s
}

// Success color style
// Success = &Theme{"success", Style{OpBold, FgGreen}}
func Success(s string) string {
	if Mode == "" {
		return color.Success.Render(s)
	}
	return s
}

// Primary color style
// Primary = &Theme{"primary", Style{OpReset, FgBlue}}
func Primary(s string) string {
	if Mode == "" {
		return color.Primary.Render(s)
	}
	return s
}

// Question color style
// Question = &Theme{"question", Style{OpReset, FgMagenta}}
func Question(s string) string {
	if Mode == "" {
		return color.Question.Render(s)
	}
	return s
}

// Secondary color style
// Secondary = &Theme{"secondary", Style{FgDarkGray}}
func Secondary(s string) string {
	if Mode == "" {
		return color.Secondary.Render(s)
	}
	return s
}

func HighlightSearchResults(text, searchPattern string) string {
	for _, pattern := range strings.Split(searchPattern, "|") {
		pattern = strings.ReplaceAll(pattern, "%", ".*?")
		compiledPattern := regexp.MustCompile(fmt.Sprintf(`(?%v)%v`, "i", pattern))
		text = compiledPattern.ReplaceAllStringFunc(text, Warn)
	}
	return text
}
