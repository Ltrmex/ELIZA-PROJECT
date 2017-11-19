package main

//Imports
import (
	"net/http"
	"fmt"
)

// Function to handle user message
func userMessage(w http.ResponseWriter, r *http.Request) {
	if (r.Method == "POST") {
		ajaxPostUserData := r.FormValue("ajaxPostUserData");	// creating and initilizing variable to user input using ajax
		fmt.Println("Received ajax data: ", ajaxPostUserData);	// printing result to the console for insurance
		
		if(ajaxPostUserData == "Maciej"){	// sample response check
			w.Write([]byte("Hello "+ ajaxPostUserData));	// sample response before implementing chat bot
		}else{
			w.Write([]byte("What is your name?"));	// alternative response
		}
	}
}
 
//Main function
func main() {
	// http.Handler
	http.Handle("/", http.FileServer(http.Dir("./static")));	// path to the static webpages
	http.HandleFunc("/receive", userMessage);	// handles userMessage function

	http.ListenAndServe(":8080", nil);	// localhost on port 8080
}