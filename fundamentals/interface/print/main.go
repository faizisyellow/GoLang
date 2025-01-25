package main

import "fmt"

type Greeting struct {
	who string
}

func (anyone Greeting) SayHi() string {
	return "Hi " + anyone.who
}

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("Hello")

	hi := Greeting{"Lizzy"}
	printSomething(hi)
	printSomething(hi.SayHi())

	printAnything(1)
	printAnything("Vortex")

	x := "lizzy mcalpine"
	printAnything(&x)

	fmt.Println(x)

}

// type interface meaning can accept all value type.
func printSomething(input interface{}) {
	// use switch type
	switch input.(type) {
	case int:
		fmt.Printf("Integer: %v\n", input)
	case float64:
		fmt.Printf("Float: %v\n", input)
	case string:
		fmt.Printf("String: %v\n", input)
	default:
		fmt.Printf("%T: %v\n", input, input)
	}

}

// using extract type from values method
func printAnything(input interface{}) {
	// when passing implicit type
	//  it will return the actual param with that type,
	//  and the second is boolean wether the param int or not.
	valueWithTypeInt, ok := input.(int)
	if ok {
		fmt.Printf("Integer: %v\n", valueWithTypeInt)
		return
	}

	valueWithTypeString, ok := input.(string)
	if ok {
		fmt.Printf("String: %v\n", valueWithTypeString)
		return
	}

	valuePointer, ok := input.(*string)
	if ok {
		*valuePointer = "hallo"
	}

}
