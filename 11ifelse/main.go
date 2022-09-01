package main

import "fmt"

func main() {
	fmt.Println("Welcome to if-else in golang")

	count := 10
	var result string

	if count < 10 {
		result = "Greater than 10"
	} else if count > 10 {
		result = "Lesser than 10"
	} else {
		result = "Equal to 10"
	}

	fmt.Println(result)

	// assigning a value and checking on the go
	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}
}
