package main

import (
	"fmt"
	"sort"
)

func find(date string, dates []string) int {
	for index, value := range dates {
		if date == value {
			return index
		}
	}

	return -1
}

func main() {
	var names, dates []string
	var points []int

	var n int
	fmt.Scan(&n)

	for n > 0 {
		var name, date string
		var point int
		fmt.Scan(&name, &point, &date)

			position := find(date, dates)
			if position == -1 {
				names = append(names, name)
				points = append(points, point)
				dates = append(dates, date)
			} else {
				if  points[position] < point {
					names[position] = name
					points[position] = point
				}
			}

		n -= 1
	}

	sort.Strings(names)

	fmt.Println(len(names))
	for _, value := range names {
		fmt.Println(value)
	}
}