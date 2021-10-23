package main

import "fmt"

func main() {
	var n int;
	fmt.Scan(&n);

	var result, count int;
	for i:= 0; i < n; i++ {
		var aux int;
		fmt.Scan(&aux);

		if (aux >= 0) {
			result += aux;
			count++;
		}
	}

	fmt.Printf("%f", float64(result) / float64(count));
}