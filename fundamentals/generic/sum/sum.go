package sum

import "fmt"

// generic function (dynamic type)
func Add[T int | float64 | string](x, y T) T {
	fmt.Printf("Passing argument x as: %T\n", x)
	fmt.Printf("Passing argument y as: %T\n", y)
	return x + y
}
