package main

import "fmt"

func main() {
	var n, h, v int;
	fmt.Scan(&n, &h, &v);

	if (n - h > n / 2) {
		h = n - h
	}

	if (n - v > n / 2) {
		v = n - v
	}

	result := h * v * 4;
	fmt.Println(result)
}