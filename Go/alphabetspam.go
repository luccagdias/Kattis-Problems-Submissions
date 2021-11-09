package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)

	inputLen := float64(len(input))
	var whitespaces, lowercases, uppercases, symbols float64
	for _, v := range input {
		if v == 95 {
			whitespaces += 1
		} else if v >= 65 && v <=90 {
			uppercases += 1
		} else if v >= 97 && v <=122 {
			lowercases += 1
		} else {
			symbols += 1
		}
	}

	fmt.Println(whitespaces / inputLen)
	fmt.Println(lowercases / inputLen)
	fmt.Println(uppercases / inputLen)
	fmt.Println(symbols / inputLen)
}