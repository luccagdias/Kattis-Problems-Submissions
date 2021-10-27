package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int64;
	fmt.Scan(&n);

	binary := strconv.FormatInt(n, 2);
	var reversedBinary string;
	for i := len(binary) - 1; i >=0 ; i-- {
		reversedBinary += string(binary[i]);
	}

	result, _ := strconv.ParseInt(reversedBinary, 2, 64);
	fmt.Println(strconv.FormatInt(result, 10));
}