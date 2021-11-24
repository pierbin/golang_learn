package main

import (
	"fmt"
)

/*
	The only difference between a slice and an array is the missing length between the brackets.
	Slices are always associated with some array, and although they can never be longer than the array, they can be smaller.

	One way to create a slice,
	y := make([]float64, 5)

	The make function also allows a third parameter: x := make([]float64, 5, 10)

	Another way to create slices is to use the [low : high] expression:
	arr := [5]float64{1,2,3,4,5}
	x := arr[0:5]

	low is the index of where to start the slice and high is the index of where to end it, but not including the index itself.

	For convenience, we are also allowed to omit low, high, or even both low and high.
	arr[0:] is the same as arr[0:len(arr)], arr[:5] is the same as arr[0:5], and
	arr[:] is the same as arr[0:len(arr)].
*/
func main() {

	arr := [6]float64{1, 2, 3, 4, 5, 6}
	x := arr[0:5]  //It means x includes the first 5 elements in arr.
	fmt.Println(x) //[1 2 3 4 5]

	appendSlice()
	copySlice()
}

/*
	append adds elements onto the end of a slice. If there’s sufficient capacity in the
	underlying array, the element is placed after the last element and the length is incremented.
	However, if there is not sufficient capacity, a new array is created, all of the
	existing elements are
*/
func appendSlice() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2) //[1 2 3] [1 2 3 4 5]
}

/*
	copy takes two arguments: dst and src. All of the entries in src are copied into dst
	overwriting whatever is there. If the lengths of the two slices are not the same, the
	smaller of the two will be used.
*/
func copySlice() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2) //It means slice2 maximum length is 2.
	copy(slice2, slice1)     //func copy(dst, src []Type) int
	fmt.Println(slice1, slice2)
}
