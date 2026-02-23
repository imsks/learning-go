package main

import "fmt"

// func greet(name string) {
// 	fmt.Printf("Hello, %s", name)
// }
˙
// func add(a int, b int) int {
// 	return a + b
// }

// func calculate(x, y int) (sum int, product int) {
// 	sum = x + yß
// 	product = x * y
// 	return
// }

type MathOperation func(int, int) int

func calculate(a, b int, op MathOperation) int {
	return op(a, b)
}

func getMultiplier(a int) func(int) int {
	return func(x int) int {
		return a * x
	}
}

func learnFunctions() {
	add := func(a, b int) int {
		return a + b
	}

	result := calculate(5, 3, add)
	fmt.Println(result)

	double := getMultiplier(2)
	fmt.Println(double(3))
}
