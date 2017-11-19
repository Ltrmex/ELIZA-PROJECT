package main

//Imports
import (
	"net/http"
)

//Main function
func main() {
	// http.Handler
	http.Handle("/", http.FileServer(http.Dir("./static")));	// path to the static webpages

	http.ListenAndServe(":8080", nil);	// localhost on port 8080
}