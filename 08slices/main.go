package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in golang")

	// slices are under the hood arrays
	var fruitList = []string{"apple", "orange", "banana"} // we don't need to give size
	fmt.Printf("Type of fruitList is %T\n", fruitList)    // []string

	// adding to slices
	fruitList = append(fruitList, "grapes", "berries", "mango")
	fmt.Println("The fruitList is: ", fruitList) // [apple orange banana grapes berries mango]

	// slicing slices
	fruitList = append(fruitList[1:6])
	fmt.Println(fruitList) // [orange banana grapes berries mango]

	fruitList = append(fruitList[:4])
	fmt.Println(fruitList) // [orange banana grapes berries]

	// using "make" keyword to initialize a slice
	highScores := make([]int, 4)
	fmt.Printf("Type of highScores is %T\n", highScores) // []int

	highScores[0] = 234
	highScores[1] = 254
	highScores[2] = 274

	fmt.Println(highScores) // [234 254 274 0]

	// slices have dynamic sizing
	highScores = append(highScores, 294, 314, 334)

	fmt.Println(highScores) // [234 254 274 0 294 314 334]

	// sorting slices (not available with arrays)
	sort.Ints(highScores)   // sorts in ascending order
	fmt.Println(highScores) // [0 234 254 274 294 314 334]

	// check is sorted
	fmt.Println(sort.IntsAreSorted(highScores)) // true

	// sort in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(highScores)))
	fmt.Println(highScores)

	// sort strings descending
	chars := []string{"a", "b", "c"}
	sort.Sort(sort.Reverse(sort.StringSlice(chars)))
	fmt.Println(chars)
	fmt.Println(len(chars))

	// remove a value from slices based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	// append can also remove values
	// x... unpacks the slice x
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
