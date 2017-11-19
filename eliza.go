package main

//Imports
import (
	"net/http"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"encoding/json"
    "io/ioutil"
	"os"
)

// Struct which will hold pattern and possible answers
type Response struct {
	pattern         *regexp.Regexp
	possibleAnswers []string
}

// Struct for loading the data from the json file
type Data struct {
	Keyword   string   `json:"keyword"`
	Responses []string `json:"responses"`
}

// Function to parse json file
func getData() []Data {
    raw, err := ioutil.ReadFile("data/responses.json")	// read json file

	if err != nil {	// error checking
        fmt.Println(err.Error())
        os.Exit(1)
    }

    var data []Data	// create array variable of struct Data to store multiple values

    json.Unmarshal(raw, &data)	//put data into the variable

    return data	// return the data
}//getData()

// Function to substitite the words e.g.you to me etc.
func substitution(input string) string {
	// Split the input on word boundaries.
	boundaries := regexp.MustCompile(`\b`)
	tokens := boundaries.Split(input, -1)
	
	// List the reflections.
	reflections := [][]string{
		{`I`, `you`},
		{`me`, `you`},
		{`you`, `me`},
		{`my`, `your`},
		{`your`, `my`},
	}
	
	// Loop through each token, reflecting it if there's a match.
	for i, token := range tokens {
		for _, reflection := range reflections {
			if matched, _ := regexp.MatchString(reflection[0], token); matched {
				tokens[i] = reflection[1]
				break
			}
		}
	}
	
	// Put the tokens back together.
	return strings.Join(tokens, ``)
}//substitution()

// Function to analyse userInput and get chatbots proper response
func analyse(userInput string) string {
	response := "";
	pages := getData()	// get the json data 
	exit := 0;	// variable used to break out of the loop
	userInput = strings.ToLower(userInput);

	for _, p := range pages {	// loop through all the data

		if (exit >= 1) {break}	// check if need to exit the loop i.e. this means response has been found

		re := regexp.MustCompile(strings.ToLower(p.Keyword)) // capture anything the after the p.Keyword
		
		for i := 0; i < len(p.Responses); i ++ {
			strings.ToLower(p.Responses[i]);
		}

		matching := Response{re, p.Responses} // variable of Response struct

		if (matching.pattern.MatchString(userInput)) {
			match := re.FindStringSubmatch(userInput)	// the pattern matched, so look for all the text after p.Keyword
			topic := match[0]	// 0 is the full string, 1 is first match.
			topic = substitution(topic)	// substitite, e.g. 'me' to 'you'
			chosenResponse := matching.possibleAnswers[rand.Intn(len(matching.possibleAnswers))]	// pick the answer
			
			// Check whenever joining strings will be required
			if(strings.Contains(p.Keyword, "(.*)")) {
				finalResponse := fmt.Sprintf(chosenResponse, topic)	// sub topic (%s) into chosenResponse
				response = finalResponse
			}else {
				response = chosenResponse	// output response
			}

			exit++	// to break out of for

		} else if(strings.Contains(userInput, "?")) {
			questions := [] string {
				"Why do you ask that?",
				"Please consider whether you can answer your own question.",
				"Perhaps the answer lies within yourself?",
				"Why don't you tell me?"}

			response = questions[rand.Intn(len(questions))]	// print random response

			exit++	// to break out of for
		}//else if
	}//for

	return response
}//analyse()

// Function to handle user message
func userMessage(w http.ResponseWriter, r *http.Request) {
	if (r.Method == "POST") {
		ajaxPostUserData := r.FormValue("ajaxPostUserData");	// creating and initilizing variable to user input using ajax
		fmt.Println("Received ajax data: ", ajaxPostUserData);	// printing result to the console for insurance
		
		w.Write([]byte(analyse(ajaxPostUserData)));
	}
}
 
//Main function
func main() {
	// http.Handler
	http.Handle("/", http.FileServer(http.Dir("./static")));	// path to the static webpages
	http.HandleFunc("/receive", userMessage);	// handles userMessage function

	http.ListenAndServe(":8080", nil);	// localhost on port 8080
}