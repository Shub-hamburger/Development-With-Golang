package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "https://lco.dev"

func main() {
	fmt.Println("Web requests in Golang")

	response, err := http.Get(URL)
	if err != nil {
		panic(err)
	}

	// if response is received
	fmt.Printf("Response type is %T\n", response) // *http.Response

	// Caller should close the connection once done (that's why we defer)
	defer response.Body.Close()

	dataBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(dataBytes)

	fmt.Println("Response is: ", content)
}
