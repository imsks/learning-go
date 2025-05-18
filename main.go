package main

import "fmt"

func main() {
	// fmt.Println("Hello World")

	// var a, b string

	// fmt.Println(a, b)
	// const a = 5
	// fmt.Println(a)

	// var i = 2

	// for i < 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// var a [5]int

	// b := [5]int{1, 2, 3, 4, 5}

	// for i, v := range b {
	// 	fmt.Println(i, v)
	// }

	// Maps
	emailSent := make(map[string]int)

	emailSent["sendgrid"] = 100
	emailSent["mailgun"] = 25

	// fmt.Println(emailSent["sendgrids"])

	for provider, count := range emailSent {
		fmt.Println(provider, '-', count)
	}
}
