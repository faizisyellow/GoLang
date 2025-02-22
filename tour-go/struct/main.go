package main

import "fmt"

type Greeting struct {
	Name, Feedback string
}

func New(name, feedback string) *Greeting {
	new := Greeting{Name: name, Feedback: feedback}

	return &new
}

func main() {
	helloTo := New("Faissal", "welcome")

	// copied (directly store the value)
	sayAgain := *helloTo

	sayAgain.Name = "Lizzy"

	fmt.Printf("%v %v\n", helloTo.Feedback, helloTo.Name)
	fmt.Printf("%v %v\n", sayAgain.Feedback, sayAgain.Name)

}
