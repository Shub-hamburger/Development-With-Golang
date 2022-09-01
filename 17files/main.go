package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file"

	// Create a file
	file, err := os.Create("./myfile.txt")
	checkNilErr(err)

	// Write in the file
	length, err := io.WriteString(file, content) // returns length or error
	checkNilErr(err)

	fmt.Println("Length is: ", length)

	// Close the file - deferred to last
	defer file.Close()

	// Read the file
	readFile("./myfile.txt")

}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	checkNilErr(err)

	fmt.Println("Text data inside file: ", string(dataByte))
}

// check for nil
func checkNilErr(err error) {
	if err != nil {
		panic(err) // panic just shuts down the execution and shows us the error
	}
}
