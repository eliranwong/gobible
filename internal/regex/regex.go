package regex

import (
	"fmt"
	"regexp"
)

// re2 syntax: https://github.com/google/re2/wiki/Syntax
// re2 does not support back reference, read my workaround at https://stackoverflow.com/a/74715496/20712208

// search and replace for general cases
func ReplaceAllString(text, flags string, searchReplace [][2]string) string {
	for _, v := range searchReplace {
		search, replace := v[0], v[1]
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

// supports regular expression search in sqlite files
// remarks: official go-sqlite example https://github.com/mattn/go-sqlite3/tree/master/_example/mod_regexp does NOT work on macOS M1
// we use custom function instead

// supports searching sqlite files with regular expression
// func Regexp(text string, pattern string, caseSensitive bool) bool
// text: source text
// pattern: regular expression pattern
// caseSensitive: determine if the search is case-sensitive
// This function is registered as a custom function.
// Therefore, users can use it in advanced bible search
// For example, to search for verses both containing words with "God" followed by "love" and containing a word "Jesus", enter:
// > .search
// > advanced
// > regexp(SCRIPTURE, "God.*?love", true) AND SCRIPTURE LIKE "%Jesus%"
// Remarks: For a quicker search, select "regexp" as search method instead of "advanced" if you have only a single pattern of regular expression.
func Regexp(text, pattern string, caseSensitive bool) bool {
	if !caseSensitive {
		pattern = fmt.Sprintf(`(?%v)%v`, "i", pattern)
	}
	return regexp.MustCompile(pattern).MatchString(text)
}

func RegexpSelect(text, pattern string, caseSensitive bool) string {
	if !caseSensitive {
		pattern = fmt.Sprintf(`(?%v)%v`, "i", pattern)
	}
	if regexp.MustCompile(pattern).MatchString(text) {
		return text
	}
	return ""
}
