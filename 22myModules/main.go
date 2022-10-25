/*
Some toolings in golang
1. go build: Builds application for production
2. go run: Runs application at a lower cost compared to go build
3. go get: Gets the assets of a library from the web
4. go env: Gives the list environment variables
5. go mod: Initializes a package
6. go mod tidy: Ensures that all imports are satisfied, removes unused libararies, imports library being used if not yet imported
7. go mod verify: Verifies the imports (using the hashes in go.sum file)
8. go list -m all: Lists the modules/libraries on which our package depends
9. go mod vendor: Constructs a directory named vendor in the main module's root directory that contains copies of all packages needed to support builds and tests of packages in the main module
*/

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

	r := mux.NewRouter()                         // create router (r)
	r.HandleFunc("/", serverHome).Methods("GET") // only serving GET method

	// Running a server
	log.Fatal(http.ListenAndServe(":4000", r)) // log.Fatal() handles failure
}

func greeter() {
	fmt.Println("Hey there!")
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	// the requests we get in inside the "r", url, params etc
	// we send response back using "w" i.e the writer
	w.Write([]byte("<h1>Welcome to Golang!</h1>"))
}
