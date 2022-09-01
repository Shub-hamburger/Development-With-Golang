package main

import "fmt"

/*
	1. structs are an alternative to classes in go
	2. go doesn't support inheritance
	3. go doesn't have "super" concept
*/
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Structs in golang")

	// using structs
	shubham := User{Name: "Shubham", Email: "shubham@gmail.com", Status: true, Age: 25}

	fmt.Println(shubham)                               // {Shubham shubham@gmail.com true 25}
	fmt.Printf("Shubham's details are %+v\n", shubham) // {Name:Shubham Email:shubham@gmail.com Status:true Age:25}

	fmt.Printf("Name is %v: and email is %v", shubham.Name, shubham.Email) // Name is shubham: and email is shubham@gmail.com
}
