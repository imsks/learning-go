package main

import "fmt"

func learnArrays() {
	// Arrays -> Passed by value
	// Fixed length. Rarely used also.

	// Slices -> Passed by Ref
	// Dynamic view.

	// Arrays
	// 1. Arrays — fixed size
	// var arr [5]int
	// arr[0] = 10
	// arr[4] = 20

	// fmt.Println(arr)

	// 2. Array literal
	// nums := [...]int{10, 20, 30, 40}
	// fmt.Println(len(nums))

	// // 3. Slice literal (no length in type)
	// s := []int{2, 3, 4}
	// s[2] = 10
	// fmt.Println(s, cap(s))

	// // append()
	// s := []int{1, 2, 3}
	// s = append(s, 4)
	// fmt.Println(s)

	// src := []int{10, 20, 30}
	// dst := make([]int, 2)
	// n := copy(dst, src)
	// fmt.Println(dst, n)

	// // Range over a slice
	// names := []string{"Alice", "Bob", "Carol"}
	// for i, name := range names {
	// 	fmt.Println(i, name) // Use _ to skip index: for _, name := range names
	// }

	// Slice from array
	arr := [5]int{1, 2, 3}
	sl := arr[1:4]
	sl[0] = 99
	fmt.Println(arr, sl)
}
