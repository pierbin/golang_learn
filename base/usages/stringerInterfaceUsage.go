package main

import "fmt"

type Stringer interface {
	String() string //Any type is fmt.Stringer, if it has a String() method that returns a string.
}

type CoffeePot string

func (c CoffeePot) String() string { //Satisfy the Stringer interface.
	return string(c) + " coffee pot" //Method needs to return a string
}

func main() {
	coffeePot := CoffeePot("LuxBrew")
	fmt.Println(coffeePot.String())
}
