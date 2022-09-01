package main

import (
	"fmt"
	"net/url"
)

const MYURL string = "http://lco.dev:3000/learn?coursename=reactjs&paymentid=dddsfdsf"

func main() {
	fmt.Println("Handling URLs in Golang")
	fmt.Println(MYURL)

	// parsing the URL
	result, err := url.Parse(MYURL)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)   // http
	fmt.Println(result.Host)     // lco.dev:3000
	fmt.Println(result.Path)     // /learn
	fmt.Println(result.Port())   // 3000
	fmt.Println(result.RawQuery) // coursename=reactjs&paymentid=dddsfdsf

	// getting query parameters
	params := result.Query()                          // returns key, value pairs
	fmt.Printf("The type of params is: %T\n", params) // url.Values

	fmt.Println(params["coursename"]) // [reactjs]
	fmt.Println(params["paymentid"])  // [dddsfdsf]

	// looping over params
	for _, val := range params {
		fmt.Println("Param is: ", val)
	}

	// creating a URL
	// we always pass a reference of URL
	partsOfUrl := &url.URL{
		Scheme:   "http",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=shubham",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
