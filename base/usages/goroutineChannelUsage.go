package main

import "fmt"

//In Go, concurrent tasks are called goroutines. Other programming languages have a similar concept called threads.
//goroutines require less memory than threads, less time to start up and stop, so it will run more goroutines at once.

//Goroutines allow for concurrency: pausing one task to work on others. And in some situations they allow parallelism.

//Go statements can’t be used with return values, because there’s no guarantee the return value will be ready
//before we attempt to use it.

//The only practical way to use a channel is to communicate from one goroutine to another goroutine.
//Not only do channels allow you to send values from one goroutine to another, they ensure the sending goroutine
//has sent the value before the receiving goroutine attempts to use it.

//Channels ensure the sending goroutine has sent the value before the receiving channel attempts to use it.
//Channels do this by blocking—by pausing all further operations in the current goroutine.
//A send operation blocks the sending goroutine until a goroutine executes a received operation on the same channel.

//To send a value on a channel, use the <- operator.  "myChannel <- 3.14", send 3.14 to myChannel.
//use the <- operator to receive values from a channel.  "<- myChannel", receive values from myChannel.

//Go programs end when the main goroutine stops, even if other goroutines have not completed their work yet.
//Send values to channels using the <- operator: myChannel <- "a value"
//The <- operator is also used to receive values from a channel: value := <-myChannel

func greeting(exampleChannel chan string, content string) { //Take a channel as a parameter
	exampleChannel <- content //Send a value to the channel
}

func firstWayCreateChannel() chan string {
	firstWayChannel := make(chan string) //Create a new channel and declare a variable at once.

	return firstWayChannel
}

func secondWayCreateChannel() chan string {
	var secondWayChannel chan string     //Declare a variable to hold a channel
	secondWayChannel = make(chan string) //Actually create the channel

	return secondWayChannel
}

func main() {
	firstWayChannel := firstWayCreateChannel()
	secondWayChannel := secondWayCreateChannel()

	go greeting(firstWayChannel, "firstWayChannel") //Pass the channel to function running in a new goroutine
	receiveFirstValue := <-firstWayChannel          //Store the received value in a variable
	fmt.Println(receiveFirstValue)                  //Receive a value from the channel

	go greeting(secondWayChannel, "secondWayChannel")
	fmt.Println(<-secondWayChannel) //Receive a value from the channel
}
