// provide shortcut utilities
package shortcuts

import (
	"fmt"
	"os"
	"path/filepath"
)

func MapKeysToSlice(inputMap map[any]any) []any {
	keys := make([]any, len(inputMap))

	i := 0
	for k := range inputMap {
		keys[i] = k
		i++
	}
	return keys
}

func MapKeysToStringSlice(inputMap map[string]string) []string {
	keys := make([]string, len(inputMap))

	i := 0
	for k := range inputMap {
		keys[i] = k
		i++
	}
	return keys
}

// walk recursively through a directory and return the paths to all files whose name matches the given pattern
// Usage:
// e.g. files, err := WalkMatch("/root/", "*.md")
func WalkMatch(root, pattern string, nameOnly bool) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if nameOnly {
		names := make([]string, len(matches))
		for i, v := range matches {
			names[i] = filepath.Base(v)[:len(filepath.Base(v))-(len(pattern)-1)]
		}
		return names, nil
	} else {
		return matches, nil
	}
}

func Divider() {
	fmt.Println("--------------------")
}
