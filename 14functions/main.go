package main

import "fmt"

// main acts as an entry point to our program
func main() {
	fmt.Println("Welcome to functions in golang")

	// calling a function
	greet("Shubham")

	// This is not allowed in golang
	// func greetTwo() {
	// 	fmt.Println("Another function")
	// }

	// This is allowed
	// In golang function literals can be used as lambda expressions
	// Anonymous functions
	func() {
		fmt.Println("Printing inside function literal")
	}()

	// assigning a function to a variable
	f := func(x int) {
		fmt.Println("current year is: ", x)
	}

	f(2021)

	result := add(2, 3)
	fmt.Println("Result is: ", result)

	result2, result3 := multiReturn()
	fmt.Printf("%v = %v\n", result2, result3)

	result4 := multiAdder(2, 3, 4)
	fmt.Println(result4)

	result5 := multiAdder(2, 3, 4, 5, 6)
	fmt.Println(result5)

}

// creating a function
func greet(name string) {
	fmt.Println("Hello, " + name)
}

// creating a function with a single return type
func add(val1 int, val2 int) int {
	return val1 + val2
}

// creating a function with multiple return types
func multiReturn() (string, int) {
	return "Two", 2
}

// creating a function with variable arguments
func multiAdder(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}

	return total
}
