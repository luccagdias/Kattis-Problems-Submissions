package main

import "fmt"

func main() {
	var cipher string
	fmt.Scan(&cipher)

	var per string
	for i := 0; i < len(cipher) / 3; i++ {
		per += "PER"
	}

	var result int
	for i, v := range cipher {
		if (string(v) != string(per[i])) {
			result += 1
		}
	}

	fmt.Println(result)
}