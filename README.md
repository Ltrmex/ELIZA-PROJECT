# Data Representation and Quering - ELIZA Project in GO

### This code was developed for college project purposes only.
### Please refer to References section at the bottom of the page as snippets of codes were implemented from external sources.

### Description
- Simple Eliza chatbot implementation in GOlang 
- Makes use of:
	- GOlang and it's various imports
	- Json file parsing and Json file handling
	- HTML - HyperText Markup Language
	- CSS - Cascading Style Sheets
	- JS - JavaScript
	- JQuery - JavaScript library designed to simplify the client-side scripting of HTML
	- Ajax - Asynchronous JavaScript and XML

I have tried to make the Eliza Chatbot code as easy to understand as possible, especially when I myself don't have lot's of 
experience in GOlang and also considering that GOlang often tends to overcomplicate things. That's the main reason why
I have chosen to store my responses data in json file as they're easier to read in. Also they allow me easier to access certain
types of data, example of that is when I use
- `re := regexp.MustCompile(strings.ToLower(p.Keyword))`
which makes use of the keyword which value of can be "My (.*)" or "I want (.*)" for example. 
	- NOTE: "(.*)" captures anything after it, for example consider sentence "Most of the time I want to relax"
	it captures "to relax" and then it does the substitutions to eventually join two strings and to reply 
	with 'What would you do if you got you want to relax?' or any other of the responses for that keyword.

To make things easier for myself I have chosen to go with easy kind of dark design for the web app which is pleasant to the eye, and
also easy for user to understand. At the first glace user is present with message from Eliza in a chat window and textbox is 
present just under the chat window, with send button on the right side of it, also
- NOTE: "Messages can be sent by just pressing enter after typing the message in textbox".

 This eye pleasing but simple design was the obvious choice for the chat bot. I was guided by the quote
- "A user interface is like a joke. If you have to explain it, it's not that good.".

###	How to run
#### NOTE: GOlang must be installed to run this code
- Step 1: Open command prompt by writing "cmd" in the search box
- Step 2: Clone repository using code below
	- `git clone https://github.com/Ltrmex/ELIZA-PROJECT`
- Step 3: Navigate from command prompt to the cloned folder location, if in the same directory e.g.
	- `cd ELIZA-PROJECT`
- Step 4: Compile the application code like so:
	- `go build eliza.go`
- Step 5: Run the code like shown below:
	- `eliza.exe`
- Step 6: Access the web application through your browser like Google Chrome or Mozzila Firefox by typing text below:
	- localhost:8080

### References
- jQuery and overall web app design was adapted from youtube tutorials:
	- [Link to jQuery, html, css tutorial!](https://www.youtube.com/playlist?list=PLHPcpp4e3JVpXbtgyD-k-nCmwpbbMIHOh)
- Transfer of data between html, jQuery and GOlan was implemented from an example below:
	- [Link to Ajax example!](https://www.socketloop.com/tutorials/golang-jquery-ajax-post-data-to-server-and-send-data-back-to-client-example)
- With never having opportunity to deal with json in GOlang I have implemented it from an example below:
	- [Link to Json file parsing example!](https://www.chazzuka.com/2015/03/load-parse-json-file-golang/)
- Chatbot implementation itself was developed from the example below:
	- [Link to Eliza chatbot simple example!](https://gist.github.com/chatton/c4328ed03eef086306ac052b89fb9797)
- Responses data and substitutions were gotten from:
	- [Link to resposnes.json data!](https://www.smallsurething.com/implementing-the-famous-eliza-chatbot-in-python/)