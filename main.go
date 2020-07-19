package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func mainRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", renderHome).Methods("GET")
	r.HandleFunc("/test", testFunc).Methods("GET")
	return r
}

func main() {
	r := mainRouter()
	http.ListenAndServe(":8000", r)
}

func testFunc(w http.ResponseWriter, r *http.Request) {
	// this will get the input from the URL path
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	fmt.Fprintf(w, "\nHello, this is just a test")
}

func renderHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("************Home Page!!!!!!")
}
