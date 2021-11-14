package main

import "fmt"

func keepOrRemove(cardsChosen int, prediction int) string {
	cardNumbers := make([]int, cardsChosen)

	for i := 0; i < cardsChosen; i++ {
		var cardNumber int
		fmt.Scan(&cardNumber)

		cardNumbers[i] = cardNumber
	}

	for _, cardNumber := range cardNumbers {
		if cardNumber == prediction {
			return "KEEP"
		}
	}

	return "REMOVE"
}

func main() {
	var numberOfCards, prediction, steps int
	fmt.Scan(&numberOfCards, &prediction, &steps)

	for i := 0; i < steps; i++ {
		var cardsChosen int
		fmt.Scan(&cardsChosen)

		keepOrRemove := keepOrRemove(cardsChosen, prediction)
		fmt.Println(keepOrRemove)
	}
}