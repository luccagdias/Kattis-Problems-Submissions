package main

import "fmt"

func main() {
	var rows, columns int
	fmt.Scan(&rows, &columns)

	parkingMap := make([]string, rows)
	for i := 0; i < rows; i++ {
		var row string
		fmt.Scan(&row)

		parkingMap[i] = row
	}

	var result [5]int
	for i := 0; i < rows - 1; i++ {
		firstRow := parkingMap[i]
		secondRow := parkingMap[i + 1]

		for j := 0; j < columns - 1; j++ {
			if (firstRow[j] == '#' || firstRow[j + 1] == '#' || secondRow[j] == '#' || secondRow[j + 1] == '#') {
				continue
			}

			var cars int
			if (firstRow[j] == 'X') {
				cars += 1
			}

			if (firstRow[j + 1] == 'X') {
				cars += 1
			}

			if (secondRow[j] == 'X') {
				cars += 1
			}

			if (secondRow[j + 1] == 'X') {
				cars += 1
			}

			result[cars] += 1
		}
	}

	for _, v := range result {
		fmt.Println(v)
	}
}