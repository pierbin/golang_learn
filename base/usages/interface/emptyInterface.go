package main

import "fmt"

/*
	So if declare an interface type that did not require any methods, it will be satisfied by any type.
	type Anything interface {}	//Can be exported anywhere.
	type anything interface {}	//Cannot be exported.
	The above interfaces are the empty interface.
	The empty interface does not require any methods to satisfy it, so it is satisfied by all types.
*/

// Do not try to call any methods on an empty interface value. If you want to call, it will have an error.
// You can try to uncomment all codes in the main and AcceptAnything method, then run it, you will see errors.

type Anything interface {
}

func AcceptAnything(thing interface{}) {
	fmt.Println(thing)
	// whistle, ok := thing.(Whistle)	//Use a type assertion to get a Whistle.
	// if ok {
	//	whistle.MakeSound()				//Call the method on the Whistle.
	// }
}

func main() {
	AcceptAnything(3.1415)
	// AcceptAnything(Whistle("Toyco Canary"))
}
