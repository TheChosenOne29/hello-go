package main

import (
	"fmt"
	"os"
)

func hola(a, b string) {
	fmt.Println("\nHello, I'm " + a + " " + b)
}
func main() {
	var name string
	var nameLast string

	fmt.Println("Hello World!")

	fmt.Print("My first name is ")
	_, err := fmt.Scan(&name)
	if err != nil {
		os.Exit(1)
	}

	fmt.Print("and my last name is ")
	_, err = fmt.Scan(&nameLast)
	if err != nil {
		os.Exit(1)
	}

	hola(name, nameLast)
}
