package main

import "fmt"

func main() {

	var number int = 10
	var decimal float64 = 3.14
	var text string = "Hello, world!"
	var isTrue bool = true

	name := "Alice"
	age := 25
	hasPet := false

	fmt.Printf("Number: %d, Type: %T\n", number, number)
	fmt.Printf("Decimal: %.2f, Type: %T\n", decimal, decimal)
	fmt.Printf("Text: %s, Type: %T\n", text, text)
	fmt.Printf("Is True: %v, Type: %T\n", isTrue, isTrue)

	fmt.Printf("Name: %s, Type: %T\n", name, name)
	fmt.Printf("Age: %d, Type: %T\n", age, age)
	fmt.Printf("Has Pet: %v, Type: %T\n", hasPet, hasPet)
}
