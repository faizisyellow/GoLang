package main

import (
	"fmt"
	"math"
)

func main() {
	// 1. Get mass input
	// 2. Squere the C. (result: 9 x 10*16)
	// 3. Multiply M = M x c*2

	var mass int64

	fmt.Print("M:")
	fmt.Scan(&mass)

	speed := int64(9 * math.Pow(10, 16))
	e := mass * speed

	fmt.Printf("E: %v", e)
}
