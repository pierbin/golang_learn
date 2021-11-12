package main

import (
	"fmt"
)

func main() {
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
