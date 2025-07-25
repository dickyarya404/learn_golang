package main

import "fmt"

func greet(name string) string {
	return "Halo, " + name + "!"
}

func main() {
	message := greet("Golang")
	fmt.Println(message)
}
