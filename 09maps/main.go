package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps in golang")

	// declaring a map
	languages := make(map[string]string)

	// adding to maps
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("Mapping of all languages: ", languages) // map[JS:Javascript PY:Python RB:Ruby]
	fmt.Println("JS stands for: ", languages["JS"])      // Javascript

	// deleting from maps
	delete(languages, "RB")
	fmt.Println("Mapping of all languages: ", languages) // map[JS:Javascript PY:Python]

	// looping over maps
	for key, value := range languages {
		// fmt.Println(key, value)
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
