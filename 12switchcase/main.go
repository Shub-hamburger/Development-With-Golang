package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Printf("Dice value is %v, roll again", diceNumber)
	case 2:
		fmt.Printf("Dice value is %v, roll again", diceNumber)
	case 3:
		fmt.Printf("Dice value is %v, roll again\n", diceNumber)
		fallthrough
	case 4:
		fmt.Printf("Dice value is %v, roll again", diceNumber)
	case 5:
		fmt.Printf("Dice value is %v, roll again", diceNumber)
	case 6:
		fmt.Printf("Dice value is %v, you can open!", diceNumber)
	default:
		fmt.Printf("Invalid input")
	}
}
