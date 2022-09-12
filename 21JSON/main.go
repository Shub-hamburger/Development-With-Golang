package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` // renames Name as "coursename" in JSON
	Price    int
	Platform string   `json:"website"`        // renames Platform as "website" in JSON
	Password string   `json:"-"`              // "-" causes password to not reflect in the JSON
	Tags     []string `json:"tags,omitempty"` // omitempty removes the fields with nil from the JSON
}

func main() {
	fmt.Println("JSON in golang")
	EncodeJSON()
	DecodeJSON()
}

func EncodeJSON() {
	courses := []course{
		{"React", 199, "Udemy", "abc", []string{"web-dev", "js"}},
		{"Vue", 299, "Coursera", "def", []string{"full-stack", "js"}},
		{"Angular", 399, "LCO", "ghi", nil},
	}

	// package the courses data as JSON
	// MarshalIndent takes -> data, prefix, indent
	finalJSON, err := json.MarshalIndent(courses, "", "\t") // we need to pass an interface (struct) into Marshaler
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJSON)
}

func DecodeJSON() {
	jsonDataFromWeb := []byte(`
		{
			"coursename": "React",
			"Price": 199,
			"Platform": "Udemy",
			"tags": ["web-dev","js"]
		}
	`)

	var courses course

	// check if jsonData is valid JSON
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON data is valid")
		// decode json
		json.Unmarshal(jsonDataFromWeb, &courses) // we pass the reference of the struct
		fmt.Printf("%#v\n", courses)
	} else {
		fmt.Printf("JSON data is not valid")
	}

	// adding data to key-value pairs
	// key will always be string, but we're not sure about the values, so we use interface
	var myData map[string]interface{}

	// decode json
	json.Unmarshal(jsonDataFromWeb, &myData)
	fmt.Printf("%#v\n", myData)

	for k, v := range myData {
		fmt.Printf("Key is %v and value is %v; type is %T\n", k, v, v)
	}
}
