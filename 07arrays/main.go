package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[2] = "Peach"
	fruitList[3] = "Apple"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var vegetablesList = [3]string{"Potato", "Beans", "Mushroom"}

	fmt.Println(vegetablesList)
}
