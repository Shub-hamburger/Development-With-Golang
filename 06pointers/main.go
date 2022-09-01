package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	var ptr *int
	fmt.Println("Value of ptr is: ", ptr) // <nil>

	myNumber := 23

	var ptr2 *int = &myNumber
	fmt.Println("ptr2 is referencing to: ", ptr2) // 0xc000012088
	fmt.Println("Value of ptr2 is: ", *ptr2)      // 23

	*ptr2 = *ptr2 + 2
	fmt.Println("New value of myNumber is: ", myNumber) // 25
}
