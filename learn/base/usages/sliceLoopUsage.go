package main

import (
	"fmt"
)

func firstWayDefineSlice() {
	var notes []string
	//When create a slice, it always needs to call make function to make the type of the slice you want to create.
	notes = make([]string, 7)

	notes[0] = "a"
	notes[1] = "b"
	fmt.Println(notes) //[a b     ]
}

func secWayDefineSlice() {
	//When create a slice, it always needs to call make function to make the type of the slice you want to create.
	primes := make([]int, 5)

	primes[1] = 2
	primes[2] = 3
	fmt.Println(primes) //[0 2 3 0 0]

	//This way, it uses slice literal to define a slice directly.
	letters := []string{"a", "b", "c"}
	fmt.Println(letters)
}

func sliceLoop() {
	//Define letters as a slice literals, the different with array literals is that empty pairs of square brackets.
	letters := []string{"a", "e", "i", "o", "u"}
	fmt.Println("The first way loop a slice.")
	for i := 0; i < len(letters); i++ {
		fmt.Println(letters[i])
	}

	fmt.Println("The second way loop a slice.")
	for _, letter := range letters {
		fmt.Println(letter)
	}
}

func main() {
	firstWayDefineSlice()
	secWayDefineSlice()
	sliceLoop()
}
