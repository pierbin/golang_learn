package main

import "fmt"

func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

/**
go run base/usages/channelBlock.go, result is: "adbecf".

It is not abcdef, because the channel block is used in this case.

Channels ensure the sending goroutine has sent the value before the receiving channel attempts to use it.
Channels do this by blocking—by pausing all further operations in the current goroutine.
A send operation blocks the sending goroutine until a goroutine executes a received operation on the same channel.

The abc goroutine blocks each time it sends a value to a channel until the main goroutine receives from it. The
def goroutine does the same. The main goroutine becomes the orchestrator of the abc and def goroutines, allowing
them to proceed only when it’s ready to read the values they’re sending.
 */
func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)
	go abc(channel1)
	fmt.Print("test")	//In go, it will run fmt.Print("test") before go channel run.
	go def(channel2)

	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Println()	//If not has the last fmt.Println(), the result will be "adbecf%". After has it, result is "adbecf"
}
