package main

import "fmt"

type Product struct {
	id    string
	title string
	price int64
}

func New(id, title string, price int64) Product {
	newProduct := Product{id: id, title: title, price: price}
	return newProduct
}

func main() {

	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	hobbies := [3]string{"watch", "coding", "soccer"}
	fmt.Println(hobbies)

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Printf("First hobbie: %v\n", hobbies[0])
	fmt.Printf("Others hobbie: %v\n", hobbies[1:])

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	hobbiesSlice := hobbies[:2]
	anotherHobbiesSlice := make([]string, 2)
	for i := range anotherHobbiesSlice {
		anotherHobbiesSlice[i] = hobbies[i]
	}
	fmt.Printf("Hobbie 1: %v\n", hobbiesSlice)
	fmt.Printf("Hobbie 2: %v\n", anotherHobbiesSlice)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	sliceAgain := hobbiesSlice[1:]
	sliceAgain = append(sliceAgain, hobbies[2])
	fmt.Printf("Another Hobbies: %v\n", sliceAgain)

	// 5) Create a "dynamic array" (slice) that contains your course goals (at least 2 goals)
	goals := []string{"understand go", "build api"}

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	goals[1] = "build rest api"
	goals = append(goals, "be backend developer")
	fmt.Printf("Goals: %v\n", goals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	products := []Product{{id: "1", title: "older", price: 33}, {id: "2", title: "five second flat", price: 75}}

	products = append(products, New("3", "march", 21))
	fmt.Println(products)
}
