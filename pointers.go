package main

func learnPointers() {
	// // 1. Basic pointer operations

	// x := 2

	// p := &x

	// fmt.Println("x =", x)
	// fmt.Println("p =", p)
	// fmt.Println("*p =", *p)

	// *p = 3
	// fmt.Println("x =", x)

	// // 2. Pointers in functions (pass by reference)
	// num := 10

	// incrementByValue(num)
	// fmt.Println("num =", num)

	// incrementByRef(&num)
	// fmt.Println("num =", num)

	// // 3. Creating pointers
	// // Method 1: Get address of existing variable
	// x := 42
	// p1 := &x

	// // Method 2: Create pointer to new value
	// p2 := new(int)
	// *p2 = 100

	// fmt.Print(p1, *p2)
}

func incrementByValue(x int) {
	x++
}

func incrementByRef(x *int) {
	*x++
}
