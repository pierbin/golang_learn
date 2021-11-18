package main

import "fmt"

/**
	An interface is a set of methods certain values are expected to have.
	Any type that has all the methods listed in an interface definition is said to satisfy that
	interface.
	A type that satisfies an interface can be assigned to any variable or function parameter
	that uses that interface as its type.

	A concrete type specifies not only what its values can do (what methods you can call on them), but also what
	they are: they specify the underlying type that holds the value’s data.

	To satisfy an interface, a type must have all the methods the interface specifies. Method names, parameter types
	(or lack thereof), and return value types (or lack thereof) all need to match those defined in the interface.

	If you’ve assigned a value of a concrete type to a variable with an interface type, you can use a type assertion
	to get the concrete type value back. Only then can you call methods that are defined on the concrete type
	(but not the interface).
*/

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

//Only methods which are defined in the interface can be called by an interface type.
//NoiseMaker interface is satisfied by any type that has a MakeSound method

type NoiseMaker interface {
	MakeSound()
}

func play(n NoiseMaker) {
	n.MakeSound()
}

func main() {
	//The comment codes are the basic way.
	//var toy NoiseMaker
	//toy = Whistle("Toyco Canary")
	//toy.MakeSound()
	//
	//toy = Horn("Toyco Blaster")
	//toy.MakeSound()

	//The below way is that declare function parameters with interface types.
	play(Whistle("Toyco Canary"))
	play(Horn("Toyco Blaster"))
}
