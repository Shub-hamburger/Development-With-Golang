package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

/*
// This is how we define a method on a struct
func (receiver Type) method_name(parameter_list) (return_type) {
	// Code
}

// Some important pointers:
1. With the help of the receiver argument, the method can access the properties of the receiver.
Here, the receiver can be of struct type or non-struct type.
2. When you create a method in your code the receiver and receiver type must be present in the same package.
*/

// Defining a method on a struct
func (user User) GetStatus() {
	fmt.Println("Is user active: ", user.Status) // Is user active:  true
}

// user here is a copy of the User struct
// so we're not modifying the original struct
func (user User) NewMail() {
	user.Email = "test@gmail.com"
	fmt.Println("Email of this user is: ", user.Email) // Email of this user is:  test@gmail.com
}

// Methods with Pointer Receiver
// We're passing a reference to the struct
// So we'll modify the original struct
func (user *User) NewAge(age int) {
	(*user).Age = age
}

// Type definition
type data int

// Defining a method on a non-struct Type
func (d1 data) multiply(d2 data) data {
	return d1 * d2
}

func main() {
	fmt.Println("Welcome to methods in golang")

	// using structs
	shubham := User{Name: "Shubham", Email: "shubham@gmail.com", Status: true, Age: 25}

	fmt.Println(shubham)                               // {Shubham shubham@gmail.com true 25}
	fmt.Printf("Shubham's details are %+v\n", shubham) // Shubham's details are {Name:Shubham Email:shubham@gmail.com Status:true Age:25}

	fmt.Printf("Name is %v: and email is %v\n", shubham.Name, shubham.Email) // Name is shubham: and email is shubham@gmail.com

	shubham.GetStatus()
	shubham.NewMail()
	fmt.Printf("Shubham's details are %+v\n", shubham) // Shubham's details are {Name:Shubham Email:shubham@gmail.com Status:true Age:25}

	// Creating a pointer
	shubham2 := &shubham

	// Passing a pointer receiver
	shubham2.NewAge(22)
	fmt.Printf("Shubham's details are %+v\n", shubham) // Shubham's details are {Name:Shubham Email:shubham@gmail.com Status:true Age:22}

	// calling method on a non struct
	value1 := data(23)
	value2 := data(20)
	res := value1.multiply(value2)
	fmt.Println("Final result is: ", res) // Final result is:  460
}
