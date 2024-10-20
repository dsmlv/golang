package main

import "fmt"

//ex1
type Person struct {
	Age  int
	Name string
}

func (p Person) Greet() {
	fmt.Println("Hello, my name is ", p.Name)
}

//ex2
type Employee struct {
	Name string
	Id   string
}

type Manager struct {
	Employee
	Department string
}

func (class Employee) Work() {
	fmt.Println(class.Id, " ", class.Name)
}

//ex3
type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func PrintArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	p := Person{
		Age:  20,
		Name: "Dinara",
	}

	p.Greet()

	manager := Manager{
		Department: "fir",
		Employee: Employee{
			Id:   "1",
			Name: "Abay",
		}}

	manager.Work()

	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 3}

	PrintArea(circle)
	PrintArea(rectangle)
}
