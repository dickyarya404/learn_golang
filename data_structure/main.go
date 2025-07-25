package main

import "fmt"

func main() {
	age := 25
	if age < 18 {
		fmt.Println("You are a minor.")
	} else if age >= 18 && age < 65 {
		fmt.Println("You are an adult.")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Current iteration:", i)
	}
}
