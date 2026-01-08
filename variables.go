package main

import "fmt"

func variables() {
	// Exercise 1: Basic Variables - Try all three declaration methods

	// Method 1: var with type (explicit)
	var name1 string
	name1 = "Sachin"

	// Method 2: var with initialization (type inferred)
	var name2 = "Go Learner"

	// Method 3: Short declaration := (most common)
	name3 := "Developer"

	fmt.Println("Exercise 1 - Variables:")
	fmt.Printf("name1: %s\n", name1)
	fmt.Printf("name2: %s\n", name2)
	fmt.Printf("name3: %s\n", name3)

	// Try declaring different types:
	var age int = 25
	var height float64 = 5.9
	var isLearning bool = true

	fmt.Printf("age: %d, height: %.1f, isLearning: %t\n", age, height, isLearning)

	// Exercise 2: Zero Values - Declare without initializing
	var zeroInt int
	var zeroFloat float64
	var zeroString string
	var zeroBool bool

	fmt.Println("\nExercise 2 - Zero Values:")
	fmt.Printf("zeroInt: %d\n", zeroInt)
	fmt.Printf("zeroFloat: %f\n", zeroFloat)
	fmt.Printf("zeroString: '%s'\n", zeroString)
	fmt.Printf("zeroBool: %t\n", zeroBool)

	// Exercise 3: Type Conversions
	var num int = 42
	var pi float64 = 3.14

	// Convert int to float64
	var numAsFloat float64 = float64(num)
	result := numAsFloat + pi

	fmt.Println("\nExercise 3 - Type Conversions:")
	fmt.Printf("num: %d\n", num)
	fmt.Printf("numAsFloat: %.2f\n", numAsFloat)
	fmt.Printf("result (numAsFloat + pi): %.2f\n", result)

	// Exercise 4: Constants
	const myName = "Sachin"
	const myAge = 25
	const piConstant = 3.14159

	fmt.Println("\nExercise 4 - Constants:")
	fmt.Printf("Name: %s, Age: %d, Pi: %.5f\n", myName, myAge, piConstant)

	// Try uncommenting this to see the error:
	// myName = "Changed"  // ERROR: cannot assign to myName

	// Exercise 5: Multiple Declarations
	var (
		firstName = "Sachin"
		lastName  = "Shukla"
		email     = "sachin@example.com"
	)

	x, y := 10, 20

	fmt.Println("\nExercise 5 - Multiple Declarations:")
	fmt.Printf("Name: %s %s, Email: %s\n", firstName, lastName, email)
	fmt.Printf("x: %d, y: %d\n", x, y)

	// Bonus: String to number conversion (requires strconv)
	// Uncomment when you want to try:
	// import "strconv" at the top
	// str := "123"
	// numFromStr, err := strconv.Atoi(str)
	// if err == nil {
	//     fmt.Printf("Converted string '%s' to int: %d\n", str, numFromStr)
	// }
}

func main() {
	variables()
}
