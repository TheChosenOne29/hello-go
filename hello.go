package main

import "fmt"

func hola(a, b string) {
	fmt.Println("\nHello, I'm " + a + " " + b)
}
func main() {
	var name string
	var nameLast string

	fmt.Println("Hello World!")

	fmt.Print("My first name is ")
	fmt.Scan(&name)

	fmt.Print("and my last name is ")
	fmt.Scan(&nameLast)

	hola(name, nameLast)
}
