package main

import (
	"fmt"
	"time"
)

func say(msg string) {
	fmt.Println(msg)
}

func learnRoutines() {
	// go say("Hi")
	// say("From Main")

	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(100 * time.Millisecond)
}
