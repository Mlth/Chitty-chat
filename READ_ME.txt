Run the server: go run server.go
Start a client: go run client.go (then write your name)
Write a message: Write your message in the console and send it
Leave the server: Terminate the program or close the console

When sending a message to the server, a clients timestamp can increase with up to 4 because it increments once when sending the message, the server increments
once when recieving the message, the server then increments once when sending back the message, and the client increments once when recieving the message again.
