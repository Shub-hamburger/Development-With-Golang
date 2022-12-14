package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

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
	return course.CourseName == ""
}

func main() {
	fmt.Println("Building API in golang")
	r := mux.NewRouter()

	// Seeding data
	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "ReactJS",
		CoursePrice: 299,
		Author:      &Author{Fullname: "Shubham", Website: "shubham.com"}})
	courses = append(courses, Course{
		CourseId:    "4",
		CourseName:  "MERN",
		CoursePrice: 199,
		Author:      &Author{Fullname: "Shubham", Website: "shubh.com"}})
	courses = append(courses, Course{
		CourseId:    "6",
		CourseName:  "MEAN",
		CoursePrice: 399,
		Author:      &Author{Fullname: "Shubham", Website: "raj.com"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listening on port 4000
	log.Fatal(http.ListenAndServe(":4000", r))

}

// controllers for different operations

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
		}
	}

	json.NewEncoder(w).Encode("No Course found")
}

// Create one course
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside")
	}

	// if course title already exists
	for _, currCourse := range courses {
		if currCourse.CourseName == course.CourseName {
			// don't add the new course
			json.NewEncoder(w).Encode("Course already exists")
		}
	}

	// generate a unique id
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	// append new course to courses slice
	courses = append(courses, course)

	json.NewEncoder(w).Encode(courses)
}

// update one course
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one course")
	w.Header().Set("Content-Type", "application/json")

	// if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// get the id from request
	params := mux.Vars(r)

	// loop over the courses and find the matching id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			// remove the found course from courses slice
			courses = append(courses[:index], courses[index+1:]...)
			// create a new course with the values from the request body
			var newCourse Course
			_ = json.NewDecoder(r.Body).Decode(&newCourse)
			// the id of this new course will remain same (as we're only updating)
			newCourse.CourseId = params["id"]
			// append the new updated course to courses slice
			courses = append(courses, newCourse)
			json.NewEncoder(w).Encode(newCourse)
		}
	}
}

// delete one course
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	// get the id from request
	params := mux.Vars(r)

	// loop over the courses and find the matching id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			// remove the found course from courses slice
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(courses)
		}
	}
}
