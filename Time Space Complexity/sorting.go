package main

import (
	"fmt"
)

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] // Tukar
			}
		}
	}
	return arr
}

// main demonstrates the use of the bubbleSort function by sorting an array of integers.
// It prints the array before and after sorting.

func main() {
	nums := []int{5, 3, 8, 2, 1}
	fmt.Println("Before:", nums)
	sorted := bubbleSort(nums)
	fmt.Println("After :", sorted)
}
