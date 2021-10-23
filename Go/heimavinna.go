package main

import (
	"fmt"
 	"strings"
	"strconv"
)

func main() {
	var input string;
	fmt.Scan(&input);

	problems := strings.Split(input, ";");
	
	var result int;
	for i:= 0; i < len(problems); i++ {
		aux := problems[i];

		if (strings.Contains(aux, "-")) {
			problemRange := strings.Split(aux, "-");
			
			num1, _ := strconv.Atoi(problemRange[0]);
			num2, _ := strconv.Atoi(problemRange[1]);

			result += num2 - num1 + 1;
		} else {
			result += 1;
		}
	}

	fmt.Println(result);
}