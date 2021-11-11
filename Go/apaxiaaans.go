package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)

	var lastChar uint8 = 0;
	for i := 0; i < len(input); i++ {
		if (input[i] != lastChar) {
			fmt.Print(string(input[i]))
		}

		lastChar = input[i]
	}
}