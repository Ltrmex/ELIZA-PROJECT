var chatLog = "";	// this variable will store previous messages

function sendMessage(message){
	// Get previous contents of the chatContainer
	chatLog = $("#chatContainer").html();

	// This is used to create new lines between chat bot and user messages
	if(chatLog.length > 4)	// check if chatLog is bigger than 4, i.e. has some content
		chatLog += "<br>";	// if it has add break line, i.e. skip to another line

	// currentMessage here is used to add more chat like feel
	$("#chatContainer").html(chatLog +"<span class='currentMessage'</span>" + "<span class='chatbot'>Eliza: </span>" + message);	// adding previous messages along with Eliza response
																																	// message and printing it to chatContainer
	$(".currentMessage").hide();	// hide currentMessage
	$(".currentMessage").delay(200).fadeIn();	// fade it in after given time
	$(".currentMessage").removeClass("currentMessage");	// remove class currentMessage
}

// First thing to show up on screen
function greeting(){
	sendMessage("Hello, I'm Eliza. How are you feeling today?");	// sends string message to sendMessage function
}

$(function(){ //	only when whole document is loaded

    greeting();	// first thing that we want to do is to print the greeting message

	$("#textbox").keypress(function(event){
		if(event.which == 13){  // 13 stands for key enter on keyboard, i.e. checks for enter input one keyboard
				$("#send").click(); // message send on button click also
				event.preventDefault(); // prevent user into going into another line when pressing enter
		}//if
	});
	
	// Here's what happens after message is sent
	$("#send").click(function(){
		var user = "<span class='username'>You: </span>";	// name tag, i.e. for chat bot it's "Eliza" and for user it's "You"

		var userMessage = $("#textbox").val();	// gets the value of user input and puts it into userMessage variable
		$("#textbox").val("");	// empties the textbox

		chatLog = $("#chatContainer").html();	// gets chatLog contents

		if(chatLog.length > 4)	// checks if anything in the chatLog
			chatLog += "<br>";	// if there is add new line

		$("#chatContainer").html(chatLog + user + userMessage);	// output the contents chatLof, user and userMessage to chatContainer, i.e. it's like overwriting charContainer 
																// contents with it's previous contents and added new message

		$("#chatContainer").scrollTop($("#chatContainer").prop("scrollHeight"));	// this line forces the chatContainer to always point to the bottom of the contents
																					// gives more like chat feeling

		// Ajax is implemented here																			
		$.ajax({
			url: 'receive',
			type: 'post',
			dataType: 'html',
			data : { 
				ajaxPostUserData: userMessage	// sends userMessage
			},
			success : function(data) {
				sendMessage(data);	// on successful response we pass in the data as parameter into sendMessage
			},
		});

	});

});
