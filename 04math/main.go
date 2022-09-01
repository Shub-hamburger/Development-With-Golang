package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcome to maths")

	var myNumberOne int = 2
	var myNumberTwo float64 = 4.5

	fmt.Println("The sum is: ", myNumberOne+int(myNumberTwo)) // 6

	// random numbers
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	// random numbers using crypto
	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNumber)
}
