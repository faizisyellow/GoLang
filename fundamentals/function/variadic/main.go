package main

import "fmt"

func main() {
	numbers := []int{3, 4, 5, 5}
	nums, feed := sum("hallo", 1, 2, 3)
	fmt.Println(nums, feed)

	// when calling a function the three dots mean spliting the slice into one each value. (splitting)
	nums1, feed1 := sum("hallo", numbers...)
	fmt.Println(nums1, feed1)
}

// params numbers(... / variadic) will be merged into a slice of the dots type. (collect all)
func sum(feedback string, numbers ...int) (int, string) {
	var result int

	for _, v := range numbers {
		result += v
	}

	return result, feedback
}
