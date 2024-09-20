package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func swap(s1, s2 string) (string, string) {
	return s2, s1
}

func quotrem(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func main() {
	sum := add(5, 3)
	fmt.Println("Sum:", sum)

	s1, s2 := swap("hello", "world")
	fmt.Println("Swapped strings:", s1, s2)

	q, r := quotrem(10, 3)
	fmt.Println("Quotient:", q, "Remainder:", r)
}
