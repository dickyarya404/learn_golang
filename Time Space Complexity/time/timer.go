package main

import (
	"fmt"
	"time"
)

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	nums := []int{5, 3, 8, 2, 1}
	fmt.Println("Before:", nums)

	start := time.Now() // Mulai pengukuran waktu

	sorted := bubbleSort(nums)

	duration := time.Since(start) // Hitung durasi eksekusi
	fmt.Println("After :", sorted)
	fmt.Println("Waktu eksekusi:", duration)
}
