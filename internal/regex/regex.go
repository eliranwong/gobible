package regex

import (
	"fmt"
	"regexp"
)

// search and replace for general cases
func ReplaceAllString(text, flags string, searchReplace [][2]string) string {
	for _, v := range searchReplace {
		search := v[0]
		replace := v[1]
		compiledPattern := regexp.MustCompile(fmt.Sprintf(`(?%v)%v`, flags, search))
		text = compiledPattern.ReplaceAllString(text, replace)
	}
	return text
}

// search and replace with loop to make sure no match is found
func ReplaceAllStringLoop(text, flags, loopPattern string, searchReplace [][2]string) string {
	compiledPattern := regexp.MustCompile(fmt.Sprintf(`(?%v)%v`, flags, loopPattern))
	for compiledPattern.MatchString(text) {
		text = ReplaceAllString(text, flags, searchReplace)
	}
	return text
}
