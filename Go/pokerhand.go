package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		ranks string
		result int
	)

	for i := 0; i < 5; i++ {
		var card string
		fmt.Scan(&card)

		ranks += string(card[0])
		aux := strings.Count(ranks, string(card[0]))

		if (aux > result) {
			result = aux
		}
	}
	
	fmt.Println(result)
}