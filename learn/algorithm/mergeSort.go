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
	fmt.Println("Sorted array is: ", mergeSort(arr))
}

func mergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	num := length / 2
	left := mergeSort(arr[:num])
	right := mergeSort(arr[num:])

	return merge(left, right)
}

func merge(left, right []int) (result []int) {
	lIndex, rIndex := 0, 0
	for lIndex < len(left) && rIndex < len(right) {
		if left[lIndex] < right[rIndex] {
			result = append(result, left[lIndex])
			lIndex++
		} else {
			result = append(result, right[rIndex])
			rIndex++
		}
	}

	result = append(result, left[lIndex:]...)
	result = append(result, right[rIndex:]...)
	return result
}
