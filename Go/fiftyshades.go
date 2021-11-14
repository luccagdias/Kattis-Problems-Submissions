package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	var result int
	for i := 0; i < n; i++ {
		var colorName string
		fmt.Scan(&colorName)

		colorName = strings.ToLower(colorName)
		if strings.Contains(colorName, "pink") || strings.Contains(colorName, "rose") {
			result += 1;
		}
	}

	if result > 0 {
		fmt.Println(result)
	} else {
		fmt.Println("I must watch Star Wars with my daughter")
	}
}