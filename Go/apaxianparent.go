package main

import (
	"fmt"
	"strings"
)	
func extendedNames(y, p string) string {
	lastLetter := string(y[len(y) - 1])
	twoLastLetters := string(y[len(y) - 2]) + string(y[len(y) - 1])

	if lastLetter == "e" {
		return y + "x" + p
	}

	if lastLetter == "a" || lastLetter == "i" || lastLetter == "o" || lastLetter == "u" {
		return strings.TrimSuffix(y, lastLetter) + "ex" + p
	}

	if twoLastLetters == "ex" {
		 return y + p
	}

	return y + "ex" + p
}

func main() {
	var y, p string
	fmt.Scan(&y, &p)

	result := extendedNames(y, p)
	fmt.Println(result)
}