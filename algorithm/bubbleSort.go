package main

import "fmt"
import "math/rand"

func main() {
	n := 7
	arr := make([]int, n)
	for i := 0; i <= n-1; i++ {
		arr[i] = rand.Intn(n)
	}

	fmt.Println("Initial array is:", arr)
	fmt.Println("")
	fmt.Println("Sorted array is: ", bubbleSort(arr))
}

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}

	return arr
}
