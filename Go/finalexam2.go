package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	answers := make([]string, n)
	for i := 0; i < n; i++ {
		var answer string
		fmt.Scan(&answer)

		answers[i] = answer
	}

	var result int
	for i := 0; i < n - 1; i++ {
		if answers[i] == answers[i + 1] {
			result += 1
		}
	}

	fmt.Println(result)
}