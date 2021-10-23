package main

import "fmt"

func main() {
	var n string;
	fmt.Scan(&n);

	result := string(n[1]) + string(n[0]);

	fmt.Println(result);
}