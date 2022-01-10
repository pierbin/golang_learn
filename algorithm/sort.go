package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n, numMax := 7, 100
	arr := generateUnsortedArr(n, numMax)
	fmt.Println("Initial array is:", arr)
	fmt.Println("Bubble sorted array is: ", BubbleSort(arr))
	fmt.Println("Quick sorted array is: ", QuickSort(arr))
	fmt.Println("Merge sorted array is: ", MergeSort(arr))
	fmt.Println("Selection sorted array is: ", SelectionQuick(arr))
	fmt.Println("Insertion sorted array is: ", insertionSort(arr))
}

func generateUnsortedArr(n int, numMax int) []int {
	rand.Seed(time.Now().Unix()) // It is used to confirm rand() will generate a random number each run time.
	arr := make([]int, n)
	for i := 0; i <= n-1; i++ {
		arr[i] = rand.Intn(int(numMax))
	}
	return arr
}

// Time complexity is O(n*n)
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}

	return arr
}

// Time complexity is O(nlogn), the best is O(nlogn), the worst is O(n*n)
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid, i := arr[0], 1
	head, tail := 0, len(arr)-1

	for head < tail {
		if arr[i] > mid {
			// This is for go special change number.
			arr[i], arr[tail] = arr[tail], arr[i]
			tail--
		} else {
			arr[i], arr[head] = arr[head], arr[i]
			head++
			i++
		}
	}

	arr[head] = mid
	QuickSort(arr[:head])
	QuickSort(arr[head+1:])

	return arr
}

func MergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	num := length / 2
	left := MergeSort(arr[:num])
	right := MergeSort(arr[num:])

	return Merge(left, right)
}

// Time complexity is O(nlogn), the best and the worst are the same, it is O(nlogn)
func Merge(left, right []int) (result []int) {
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

// Time complexity is O(n*n), the best is O(n*n), the worst is O(n*n)
func SelectionQuick(arr []int) []int {
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
			// This is for go special change number.
			arr[i], arr[min] = arr[min], arr[i]
		}
	}

	return arr
}

// Time complexity is O(n*n), the best is O(n), the worst is O(n*n)
func insertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}
