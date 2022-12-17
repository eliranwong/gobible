// Package check
// check information or errors
package check

import (
	"log"
	"os"
)

func DbErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// a generic function to check if an element is in a slice
func SliceContainsElement[T comparable](slice []T, element T) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}
