package main

import "fmt"

func solve(input string) string {
	wordLength := len(input) / 3

	firstWord := input[:wordLength]
	secondWord := input[wordLength : 2*wordLength]
	thirdWord := input[2*wordLength : 3*wordLength]

	if (firstWord == secondWord || firstWord == thirdWord) {
		return firstWord
	}

	return secondWord
}
func main() {
	var input string
	fmt.Scan(&input)

	fmt.Println(solve(input))
}