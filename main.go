package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// tutorial source: https://tutorialedge.net/golang/creating-restful-api-with-golang/
// code from here: https://github.com/TutorialEdge/create-rest-api-in-go-tutorial/blob/master/main.go

// To get started we will have to create a very simple server which can handle HTTP requests.
// To do this we’ll create a new file called main.go. Within this main.go file we’ll want to define 3 distinct
// functions. A homePage function that will handle all requests to our root URL, a handleRequests function
// that will match the URL path hit with a defined function and a main function which will kick off our API.

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
