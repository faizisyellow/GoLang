package main

import (
	"fmt"
	"sort"
)

func main() {
	slice1 := make([]int, 3, 3)
	sliceAddr1 := &slice1
	fmt.Printf("Slice variable address: %p\n", sliceAddr1)

	slice1 = append(slice1, 4)
	sliceAddr2 := &slice1
	fmt.Printf("Slice variable address: %p, add: %p\n", sliceAddr2, slice1)
	// Both addresses will be the same

	numbers := []int{3, 2, 4, 1}

	sort.Ints(numbers)

	fmt.Println(numbers)
}
