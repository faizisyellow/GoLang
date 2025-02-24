package main

import "fmt"

type Address struct {
	street, city string
}

type Greeting struct {
	Address
	Name, Feedback string
}

func New(name, feedback, address, city string) *Greeting {
	new := Greeting{Name: name, Feedback: feedback,
		Address: Address{
			street: address,
			city:   city,
		},
	}

	return &new
}

func main() {
	helloTo := New("Faissal", "welcome", "10th lincoln", "new york")

	// copied (directly store the value)
	sayAgain := *helloTo

	sayAgain.Name = "Lizzy"

	fmt.Printf("%v %v\n", helloTo.Feedback, helloTo.Name)
	fmt.Printf("%v %v\n", sayAgain.Feedback, sayAgain.Name)
	fmt.Println((*helloTo).Address.city)
}
