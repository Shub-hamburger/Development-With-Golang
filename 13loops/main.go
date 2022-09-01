package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

	// for loop
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// for loop variety 2
	for i := range days {
		fmt.Println(days[i])
	}

	// kind of for each loop
	for index, day := range days {
		fmt.Println(index, day)
	}

	// kind of while loop with continue, break and goto
	val := 1
	for val < 10 {
		if val == 6 {
			goto lco
		}
		if val == 2 {
			val++
			continue
		}
		if val == 8 {
			break
		}

		fmt.Println("Value is: ", val)
		val++
	}

lco: // goto location
	fmt.Println("This is my location")
}
