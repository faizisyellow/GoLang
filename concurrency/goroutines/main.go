package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, done chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)

	done <- true
}

func main() {
	done := make(chan bool)
	// greet("Nice to meet you!")
	// greet("How are you?")
	greet("I hope you're liking the course!")
	go slowGreet("How ... are ... you ...?", done)
	<-done
}
