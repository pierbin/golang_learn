package main

import (
	"bufio" //When read the input from keyboard, will be used.
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

//Using greeting() to get a grade.
func getFloat() (float64, error) {
	//read (receive and store) input from the program’s standard input, which all keyboard input goes to
	//NewReader will return a new bufio.Reader
	//os.Stdin store the standard input from the keyboard.
	reader := bufio.NewReader(os.Stdin)

	//Return everything the user has typed, up to where they pressed the Enter key
	//ReadString will return the user typed, as a string. The ReadString method requires an argument with a rune (character) that marks the end of the input.
	//'\n' means everything up until the newline rune will be read.
	// It should not be named error, because it is the go keyword. That's why named err.
	input, err := reader.ReadString('\n') //err can be replaced by '_', then the error will be ignored.
	if err != nil {
		//Report the error and stop the program
		// log.Fatal(err)
		return 0, err
	}

	fmt.Println(reflect.TypeOf(input), input)

	//TrimSpace is used to remove all whitespace characters (newlines, tabs, and regular spaces) from the start and end of a string. from the input string
	input = strings.TrimSpace(input)

	//Convert the string to a float64 value
	//ParseFloat is used to convert the string to a number, and returns it as a float64 value.

	/*
		When the same variable name is declared twice in the same scope, we will get a compile error.
		As long as at least one variable name in a short variable declaration is new, it’s allowed.
	*/
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	return number, nil
}

func main() {
	//Print doesn’t skip to a new terminal line after printing a message, which lets us keep the prompt and the user’s entry on the same line
	fmt.Print("Enter a grade: ")

	grade, err := getFloat()
	if err != nil {
		log.Fatal(err)
	}

	//If not declare the status outside the condition, it will have a variable scope error.
	//In go, the variable scope is different with PHP.
	//It has condition scope, function scope, package scope and file scope. So if a variable wants to use outside a condition scope, it must be declared outside the condition scope first.
	//condition scope is also called condition block. Then function scope equals to function block. Others are the same.
	var status string
	if grade >= 60 {
		status = "You passed"
	} else {
		status = "You failed"
	}

	fmt.Println(reflect.TypeOf(grade), grade)
	fmt.Println(status)
}
