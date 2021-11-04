package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var input string
		fmt.Scan(&input)

		if (string(input[0]) == "P") {
			fmt.Println("skipped")
		} else {
			nums := strings.Split(input, "+")

			num1, _ := strconv.Atoi(nums[0])
			num2, _ := strconv.Atoi(nums[1])
			fmt.Println(num1 + num2)
		}
	}
}