package main

import "fmt"

func main() {
	var r1, s int;
	fmt.Scan(&r1, &s);

	result := (2 * s) - r1;
	fmt.Println(result);
}