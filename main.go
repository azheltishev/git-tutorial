package main

import "fmt"

func main() {
	var x string

	fmt.Println("What is your name?")
	fmt.Scan(&x)

	fmt.Println("Hello", x, "!")
}
