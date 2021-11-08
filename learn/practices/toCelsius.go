package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

/*
	If parts of your code are shared between multiple programs, you should consider moving them into greeting.
	The greeting() method is used in this file and inputGrade.go file, so it should be moved into a package.

	Receive the input from command line and convert it to float.
	Return float64 and error.
*/
func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return 0, err
	}

	fmt.Println(reflect.TypeOf(input), input)

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}

	return number, nil
}

func main() {
	fmt.Println("Enter a temperature in Fahrenheit:")

	//Call greeting() to get a temperature.
	fahrenheit, err := getFloat()

	if err != nil {
		log.Fatal(err)
	}

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celsius\n", celsius)
}
