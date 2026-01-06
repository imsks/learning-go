package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	// Variables and Constants
	// 1. var with type (explicit)
	var name string
	name = "Sachin"
	fmt.Println(name)

	// 2. var with type (implicit)
	var age int = 20
	fmt.Println(age)

	// 3. var with type (implicit)
	var isStudent bool = true
	fmt.Println(isStudent)

	// 4. Shorthand
	name1 := "Sachin"
	fmt.Println(name1)

	var x, y int = 1, 2
	fmt.Println(x, y)
	a, b := 3, 4
	fmt.Println(a, b)
	var (
		name2 string = "Sachin"
		age2  int    = 20
	)
	fmt.Println(name2, age2)

	const pi = 3.14159
	const (
		StatusOK    = 200
		StatusError = 500
	)

	fmt.Println(pi, StatusOK, StatusError)
}
