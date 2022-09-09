package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("GET & POST Requests in Golang")
	PeformGetRequest()
	PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PeformGetRequest() {
	const MYURL = "http://localhost:1111/get"

	response, err := http.Get(MYURL)
	if err != nil {
		panic(err)
	}

	// close the request after done processing
	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	// Builders are used to efficiently build a string using Write methods. It minimizes memory copying
	var responseString strings.Builder

	// read response
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String()) // responseString has the response now

	fmt.Println(string(content)) // returns the same content
}

func PerformPostJsonRequest() {
	const MYURL = "http://localhost:1111/post"

	// json payload
	requestBody := strings.NewReader(`
		{
			"coursename":"Learn golang",
			"price":0,
			"platform":"YouTube"
		}
	`)

	response, err := http.Post(MYURL, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	// close the request once done
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const MYURL = "http://localhost:1111/postform"

	// create form data
	data := url.Values{}
	data.Add("firstname", "Shubham")
	data.Add("lastname", "Raj")
	data.Add("email", "shubham@go.dev")

	response, err := http.PostForm(MYURL, data)
	if err != nil {
		panic(err)
	}

	// close request when done
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}
