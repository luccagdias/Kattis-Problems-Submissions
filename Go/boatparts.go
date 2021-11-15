package main

import "fmt"

func hasPart(boatParts []string, part string) bool {
	for _, value := range boatParts {
		if value == part {
			return true
		}
	}

	return false
}

func main() {
	var p, n, partsReplaced int
	fmt.Scan(&p, &n)

	boatParts := make([]string, n)
	for i := 0; i < n; i++ {
		var part string
		fmt.Scan(&part)

		hasPart := hasPart(boatParts, part)
		if !hasPart {
			boatParts[i] = part
			partsReplaced += 1

			if partsReplaced == p {
				fmt.Println(i + 1)
				break
			}
		}
	}

	if partsReplaced != p {
		fmt.Println("paradox avoided")
	}

}