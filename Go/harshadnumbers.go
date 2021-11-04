package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	for {
		str := strconv.Itoa(n)

		totalSum := 0
		for i := 0; i < len(str); i++ {
			digit, _ := strconv.Atoi(string(str[i]))
			totalSum += digit 
		}

		if n % totalSum == 0 {
			fmt.Println(n)
			break;
		}

		n += 1
	} 
}