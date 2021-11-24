package main

import "fmt"
import "math/rand"

func main() {
	n := 10
	arr := make([]int, n)
	for i := 0; i <= n-1; i++ {
		arr[i] = rand.Intn(n)
	}

	fmt.Println("Initial array is:", arr)
	fmt.Println("")
	fmt.Println("Sorted array is: ", quickSort(arr))
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid, i := arr[0], 1
	head, tail := 0, len(arr)-1

	for head < tail {
		if arr[i] > mid {
			//This is for go special change number.
			arr[i], arr[tail] = arr[tail], arr[i]
			tail--
		} else {
			arr[i], arr[head] = arr[head], arr[i]
			head++
			i++
		}
	}

	arr[head] = mid
	quickSort(arr[:head])
	quickSort(arr[head+1:])

	return arr
}
