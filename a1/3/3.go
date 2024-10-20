package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	if a == 0 {
		fmt.Print("zero")
	} else if a < 0 {
		fmt.Println("negative")
	} else {
		fmt.Println("positive")
	}

	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}

	fmt.Println("The sum of first 10 natural numbers is: ", sum)

	// sum = 0
	// for index := range 10 {
	// 	sum += index
	// }
	// fmt.Println(sum)

	var d int
	fmt.Println("Enter a day 1 to 7: ")
	fmt.Scan(&d)

	switch d {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("more than 7")
	}
}
