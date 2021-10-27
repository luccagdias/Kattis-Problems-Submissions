package main

import "fmt"

func main() {
	var input string;
	fmt.Scan(&input);

	var Ts, Cs, Gs int;
	for i := 0; i < len(input); i++ {
		aux := string(input[i]); 
		
		if aux == "T" {
			Ts += 1;
		}

		if aux == "C" {
			Cs += 1;
		}

		if aux == "G" {
			Gs += 1;
		}
	}

	min := Ts;
	if (min > Cs) {
		min = Cs;
	}

	if (min > Gs) {
		min = Gs;
	}

	result := (Ts * Ts) + (Cs * Cs) + (Gs * Gs) + min * 7;

	fmt.Println(result);
}