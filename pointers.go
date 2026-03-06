package main

// type Person struct {
// 	Name string
// 	Age  int
// }

// func (p Person) SetAgeByValue(age int) {
// 	p.Age = age
// }

// func (p *Person) SetAgeByRef(age int) {
// 	p.Age = age
// }

func learnPointers() {
	// // 1. Basic pointer operations

	// x := 2

	// p := &x // Referenced

	// fmt.Println("x =", x)
	// fmt.Println("p =", p)
	// fmt.Println("*p =", *p) // De-Referenced

	// *p = 3
	// fmt.Println("x =", x)

	// // 2. Pointers in functions (pass by reference)
	// num := 10

	// incrementByValue(num)
	// fmt.Println("num =", num)

	// incrementByRef(&num)
	// fmt.Println("num =", num)

	// // // 3. Creating pointers
	// // // Method 1: Get address of existing variable
	// x := 42
	// p1 := &x

	// // // Method 2: Create pointer to new value
	// p2 := new(int) // allocates memory for an int and returns a pointer
	// *p2 = 100      // Assigns value to allocated memory

	// fmt.Print(p1, p2, *p2)

	// var p *int // Nil pointer

	// // fmt.Println(p, *p) // Error: Can't dereference a Nill pointer

	// if p != nil {
	// 	fmt.Println(*p)
	// } else {
	// 	fmt.Println("Pointer is nil")
	// }

	// person := Person{Name: "Sachin", Age: 25}

	// person.SetAgeByValue(30)
	// fmt.Println(person.Age)

	// person.SetAgeByRef(35)
	// fmt.Println(person.Age)

	// When to use Pointers vs Values
	// Use VALUES when:
	// - Small data (int, bool, string)
	// - You don't need to modify original
	// - You want immutability
	// func add(a, b int) int {
	// 	return a + b
	// }

	// Use POINTERS when:
	// - Large structs (avoid copying)
	// - Need to modify original
	// - Optional values (nil = not set)
	// func updatePerson(p *Person) {
	// 	p.Age++
	// }

}

// // Exercise 1: Write a function that swaps two integers using pointers
// func swap(a, b *int) {
// 	temp := &a
// 	b = a
// 	a = *temp
// 	fmt.Println(a, b)
// }

// func incrementByValue(x int) {
// 	x++
// }

// func incrementByRef(x *int) {
// 	*x++
// }
