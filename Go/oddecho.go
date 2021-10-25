package main

import "fmt"

func main() {
	var n int;
	fmt.Scan(&n);

	for i := 1; i <= n; i++ {
		var word1 string;
		fmt.Scan(&word1);

		if (i % 2 != 0) {
			fmt.Println(word1);
		}
	}
}