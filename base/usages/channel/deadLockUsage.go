package main

import "fmt"

func main() {
	// Use the make keyword to create a new object. Specify what type of data is returned from the channel using string.
	dataChannel := make(chan string)

	// Telling the main function to keep waiting until the channel receives the data.
	fmt.Println(<-dataChannel)

	/**
	In the program above, we don’t have any other goroutines sending data to the channel.
	Since there is no other channel available, the Golang program will encounter a deadlock waiting to receive the data.
	*/
}

/**
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
        /Users/jiangdawei/go/src/learnGo/base/usages/channel/deadLockUsage.go:6 +0x36
exit status 2
*/
