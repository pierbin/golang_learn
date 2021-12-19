package main

import "fmt"

// Adding a buffer will provide storage capacity to the channel, and it need not be disposed of immediately.
// Since it’s buffered, the main channel will move further when we’re adding more data to the channel.
// Once the buffer is full, we need to dispose the data or else the deadlock will be achieved again.
func main() {
	dataChannel := make(chan string, 3) // Creating a channel with a buffer.
	dataChannel <- "Some Sample Data"
	dataChannel <- "Some Other Sample Data"
	dataChannel <- "Buffered Channel"

	// After buffer a channel, output is from the first buffer to the last one.
	fmt.Println(<-dataChannel)
	fmt.Println(<-dataChannel)
	fmt.Println(<-dataChannel)
}
