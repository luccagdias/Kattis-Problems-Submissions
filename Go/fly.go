package main

import "fmt"

func main() {
	var phone string;
	fmt.Scan(&phone);

	result := 1;
	for i := 0; i < 3; i++ {
		if (phone[i] != '5') {
			result = 0;
		}
	}

	fmt.Println(result);
}