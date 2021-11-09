package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scan(&input)

	var words = strings.Split(input, "-")
	for _, v := range words {
		fmt.Print(string(v[0]))
	}

	fmt.Println()
}