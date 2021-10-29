package main

import "fmt"

func bestVictoryCard(buyingPower int) {
	var output string

	if buyingPower >= 8 {
		output = "Province or "
	} else if buyingPower >=5 {
		output ="Duchy or "
	} else if buyingPower >= 2 {
		output ="Estate or "
	}

	fmt.Print(output)
}

func bestTreasureCard(buyingPower int) {
	var output string

	if buyingPower >= 6 {
		output = "Gold "
	} else if buyingPower >=3 {
		output ="Silver"
	} else if buyingPower >= 0 {
		output ="Copper"
	}

	fmt.Println(output)
}

func main() {
	var g, s, c int
	fmt.Scan(&g, &s, &c)

	buyingPower := (g * 3) + (s * 2) + c
	bestVictoryCard(buyingPower)
	bestTreasureCard(buyingPower)
}