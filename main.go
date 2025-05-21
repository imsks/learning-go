package main

import "fmt"

func main() {
	type EmailRequest struct {
		To      string
		Subject string
		HTML    string
	}

	email := EmailRequest{
		To:      "hello@tbe.in",
		Subject: "Welcome to Chitthi",
		HTML:    "<h1>Hello World</h1>",
	}

	fmt.Println(email.HTML)

}

// func main() {
// 	// fmt.Println("Hello World")

// 	// var a, b string

// 	// fmt.Println(a, b)
// 	// const a = 5
// 	// fmt.Println(a)

// 	// var i = 2

// 	// for i < 5 {
// 	// 	fmt.Println(i)
// 	// 	i++
// 	// }

// 	// var a [5]int

// 	// b := [5]int{1, 2, 3, 4, 5}

// 	// for i, v := range b {
// 	// 	fmt.Println(i, v)
// 	// }

// 	// Maps
// 	// emailSent := make(map[string]int)

// 	// emailSent["sendgrid"] = 100
// 	// emailSent["mailgun"] = 25

// 	// // fmt.Println(emailSent["sendgrids"])

// 	// for provider, count := range emailSent {
// 	// 	fmt.Println(provider, '-', count)
// 	// }

// 	// fmt.Println(sum(1, 2))

// 	result, err := divide(10, 2)

// 	if err != nil {
// 		fmt.Println("Error: ", err)
// 	} else {
// 		fmt.Println("Result: ", result)
// 	}

// 	square := func(x int) int {
// 		return x * x
// 	}

// 	fmt.Println(square(2))
// }

// // Functions
// // func sayHello() {
// // 	fmt.Println("Hello")
// // }

// // func sum(a int, b int) int {
// // 	return a + b
// // }

// func divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("cannot divide by zero")
// 	}

// 	return a / b, nil
// }
