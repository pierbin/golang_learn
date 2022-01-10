package main

import "fmt"

func main() {
	// Prompt the user to enter a number
	fmt.Print("Enter a number: ")

	var input float64
	a := &input

	// Scanf fills input with the number we enter
	fmt.Scanf("%f", a)
	output := input * 2
	fmt.Println(output)
}
