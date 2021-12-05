package main

import (
	"fmt"
)

func firstWayDefineSlice() {
	var notes []string
	//When create a slice, it always needs to call make function to make the type of the slice you want to create.
	//The number 7 means the slice length is 7.
	notes = make([]string, 7)

	notes[0] = "a"
	notes[1] = "b"
	fmt.Println(notes) //[a b     ]

	letters := make([]int, 5, 10)
	var testNil []int                                    //An empty slice default value is nil, so it is a nil slice, it won't have any elements in it.
	fmt.Println(letters == nil, testNil == nil, letters) //false true [0 0 0 0 0] All default is 0.
}

func secWayDefineSlice() {
	//When create a slice, it always needs to call make function to make the type of the slice you want to create.
	//The number 5 means the slice length is 5.
	primes := make([]int, 5)

	primes[1] = 2
	primes[2] = 3
	fmt.Println(primes) //[0 2 3 0 0]

	//This way, it uses slice literal to define a slice directly.
	letters := []string{"a", "b", "c"}
	fmt.Println(letters) //[a b c]
}

func sliceLoop() {
	//Define letters as a slice literals, the different with array literals is that empty pairs of square brackets.
	letters := []string{"a", "e", "i", "o", "u"}
	fmt.Println("The first way loop a slice.")
	for i := 0; i < len(letters); i++ {
		fmt.Printf("%v ", letters[i])
	}
	fmt.Println("")

	fmt.Println("The second way loop a slice.")
	for _, letter := range letters {
		fmt.Printf("%v ", letter)
	}
	fmt.Println("")
}

func getSliceElements() {
	letters := []string{"a", "b", "c", "d", "e"}

	//0 means index of array where slice should start,
	//3 means index of array slice should stop before.
	//The slice1 elements will include 0<=index<3.
	slice1 := letters[0:3] //[a b c]

	//The slice operator always start default is 0, if omit the start index, 0 will be used.
	slice2 := letters[:2] //[a b]

	//If the stop index is omitted, it will stop at the last element of an array.
	slice3 := letters[2:] //[c d e]

	//There is no index 5, but the array will stop at the last element.
	//Must make sure, the index won't go further than the len(array), otherwise it will have an error.
	slice4 := letters[3:5] //[d e]

	fmt.Println(slice1, slice2, slice3, slice4)
}

func appendUsage() {
	//make([]int, 5, 10) allocates an underlying array of size 10 and returns a slice of length 5 and capacity 10.
	//The slice has an initial length of 5 and has a capacity of 10.
	//The capacity means the slice letters maximum capacity is 10.
	//If add elements over capacity and not use append to add, it will have an error.
	letters := make([]int, 5, 10)
	fmt.Printf("letters capacity is %v, length is %v, content is %v.\n", cap(letters), len(letters), letters)

	letters = append(letters, 6)
	behindSlice := []int{7, 8}
	//If append a slice into the other slice, the end of merge slice must have "..."
	letters = append(letters, behindSlice...)
	fmt.Printf("letters capacity is %v, length is %v, content is %v.\n", cap(letters), len(letters), letters)

	frontSlice := []int{-1, -2}
	letters = append(frontSlice, letters...) //insert slice in the front of a slice.
	fmt.Printf("letters capacity is %v, length is %v, content is %v.\n", cap(letters), len(letters), letters)

	//If the next two lines won't comment, it will have an error.
	//Because set value directly to letters, it is a static way to add element into a slice.
	//But using append() method to add elements into a slice, it is a dynamic way to increase elements for a slice.
	//letters[8] = 10
	//fmt.Println(letters)
	letters = append(letters, 11, 12, 13)
	fmt.Printf("letters capacity is %v, length is %v, content is %v.\n", cap(letters), len(letters), letters)
}

//When a copy happens, it always replace elements from the start position.
func copyUsage() {
	var s1 []int         //Define a nil slice. Nothing can be copied into a nil slice.
	s2 := []int{1, 2, 3} //During a copy happens, it won't over the capacity of the first slice.
	s3 := []int{4, 5, 6, 7, 8}
	s4 := []int{9, 10}

	n1 := copy(s1, s2)
	fmt.Printf("n1=%d, s1=%v, s2=%v\n", n1, s1, s2) //n1=0, s1=[], s2=[1 2 3]

	n2 := copy(s2, s3)
	fmt.Printf("n2=%d, s2=%v, s3=%v\n", n2, s2, s3) //n2=3, s2=[4 5 6], s3=[4 5 6 7 8]

	n3 := copy(s3, s4)
	fmt.Printf("n3=%d, s3=%v, s4=%v\n", n3, s3, s4) //n3=2, s3=[9 10 6 7 8], s4=[9 10]
}

func main() {
	firstWayDefineSlice()
	secWayDefineSlice()
	sliceLoop()
	getSliceElements()
	appendUsage()
	copyUsage()
}
