package main

import "fmt"

/*
	1. Statements starting with defer are executed at the end
	2. The defer statements are executed in LIFO fashion i.e. the last defer statement is executed first
*/

//World! 4 3 2 1 0 Hello Yo
func main() {
	defer fmt.Print("Yo ")
	defer fmt.Print("Hello ")
	fmt.Print("World! ")

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%v ", i)
	}
}
