package main

import (
	"fmt"
	"strings"
)

func main() {
	var word, permutation string
	fmt.Scan(&word, &permutation)

	var components int
	remainingLetters := len(word)

	for _, letter := range permutation {
		occurences := strings.Count(word, string(letter))
		if occurences > 0 {
			remainingLetters -= occurences
		} else {
			components += 1
		}


		if components == 10 {
			fmt.Println("LOSE")
			break
		}

		if remainingLetters == 0 {
			fmt.Println("WIN")
			break
		}
	}
}