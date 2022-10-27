package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Model for course
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"` // injecting Author in Course
	// we pass reference of Author as we don't want a copy of Author, we want a Type of Author
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// data for API
var courses []Course

// middleware
func (course *Course) IsEmpty() bool {
	return course.CourseId == "" && course.CourseName == ""
}

func main() {
	fmt.Println("Building API in golang")

}

// controllers

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to our API</h1>"))
}

// GET all courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET all courses")

	// setting headers
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(courses) // throws back the data given to Encode in JSON format back to the caller via ResponseWriter
}

// GET one course
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET one course")
	// grab id of course from the request
	params := mux.Vars(r)
	fmt.Printf("Params are %v", params)
	// setting headers
	w.Header().Set("Content-Type", "application/json")

	// loop through courses, find the matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course found")
	return
}
