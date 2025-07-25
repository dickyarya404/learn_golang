package main

import "fmt"

func main() {
	arr := [7]int{1, 2, 7, 4, 5, 12, 32}
	slice := []int{1, 2, 7, 4, 5, 12, 32} // Mengambil elemen dari indeks 1 hingga

	slice = append(slice, 5) // Menambahkan elemen baru ke slice

	fmt.Println("Array:", arr)
	fmt.Println("Slice:", slice)

}
