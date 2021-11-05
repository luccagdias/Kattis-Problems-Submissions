package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int;
	fmt.Scan(&n);

	for i := 0; i < n; i++ {
		var line1, line2 string
		fmt.Scan(&line1, &line2)

		var result strings.Builder
		for i, value := range line1 {
			if string(value) == string(line2[i]) {
				result.WriteString(".")
			} else {
				result.WriteString("*")
			}
		}

		fmt.Println(line1)
		fmt.Println(line2)
		fmt.Println(result.String())
		fmt.Println()
	}
}