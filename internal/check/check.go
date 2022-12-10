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

func StringInSlice(element string, slice []string) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func IntInSlice(element int, slice []int) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}
