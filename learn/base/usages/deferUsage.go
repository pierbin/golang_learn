package main

import (
	"fmt"
	"log"
)

/**
	You can place defer keyword before any ordinary function or method call, and Go will defer (that is, delay)
	making the function call until after the current function exits.

	The "defer" keyword ensures a function call takes place even if the calling function exits early by using the
	return keyword.

	The "defer" keyword can only be used with a function or method call, it cannot be used in for loops or assignments.
 */

func Socialize() error {
	defer fmt.Println("Test defer")

	fmt.Println("Hello")
	return fmt.Errorf("Test return error after defer, what will happen.")	//return an error

	//The below codes won't run.
	fmt.Println("The middle not the last output, after the defer.")
	return nil
}

/**
	The output is as below
	Hello
	The middle not the last output, after defer.
	Test defer.	//The first output is deferred until after Socialize() exits.
 */
func main() {
	err := Socialize()
	if err != nil {
		log.Fatal(err)
	}
}
