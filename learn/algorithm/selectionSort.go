package main

import "fmt"
import "math/rand"

func main() {
	n := 5
	arr := make([]int, n)
	for i := 0; i <= n-1; i++ {
		arr[i] = rand.Intn(n)
	}

	fmt.Println("Initial array is:", arr)
	fmt.Println("")
	fmt.Println("Sorted array is: ", selectionQuick(arr))
}

func selectionQuick(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	min := 0

	for i := 0; i < len(arr); i++ {
		min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		if min != i {
			//This is for go special change number.
			arr[i], arr[min] = arr[min], arr[i]
		}
	}

	return arr
}
