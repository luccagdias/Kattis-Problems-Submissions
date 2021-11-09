package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	lastDigit := 0
	hasMissingNumber := false
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)

		if (lastDigit != num - 1) {
			hasMissingNumber = true;
			for i := lastDigit + 1; i < num; i++ {
				fmt.Println(i)
			}
		}

		lastDigit = num
	}

	if !hasMissingNumber {
		fmt.Println("good job")
	}
}