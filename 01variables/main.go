package main

import "fmt"

// constants - cannot be changed
// First letter should be capitalized as this
// variable is "public"
const LoginToken string = "ddsdfdsf"

func main() {
	fmt.Println(LoginToken)
	fmt.Printf("LoginToken is of type: %T \n", LoginToken)

	var username string = "Shubham"
	fmt.Println(username)
	fmt.Printf("username is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("isLoggedIn is of type: %T \n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("smallValue is of type: %T \n", smallValue)

	var smallFloat float32 = 255.7668756585977
	fmt.Println(smallFloat)
	fmt.Printf("smallFloat is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("anotherVariable is of type: %T \n", anotherVariable)

	// implicit declaration
	var company = "amazon"
	fmt.Println(company)

	// no var declaration using walrus operator
	// can only be used inside a method
	numberOfUsers := 30000
	fmt.Println(numberOfUsers)
	fmt.Printf("numberOfUsers is of type: %T \n", numberOfUsers)
}
