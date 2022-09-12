package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Mod in golang")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET") // only serving GET method

	// Running a server
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey there!")
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	// we send response back using w i.e the writer
	w.Write([]byte("<h1>Welcome to Golang!</h1>"))
}
