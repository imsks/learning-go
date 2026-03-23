package main

import (
	"fmt"
	"os"
)

// func divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, errors.New("division by zero")
// 	}
// 	return a / b, nil
// }

func learnErrorHandling() {
	// // caller:
	// result, err := divide(10, 0)
	// if err != nil {
	// 	fmt.Println("failed:", err)
	// 	return
	// }
	// fmt.Println(result)
	_, err := os.ReadFile("file.txt")
	if err != nil {
		fmt.Println("failed:", err)
		return
	}
	fmt.Println(err)
}
