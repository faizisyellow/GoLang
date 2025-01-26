package main

import (
	"fmt"

	"example.com/generic/min"
	"example.com/generic/operation"
	"example.com/generic/sum"
)

// constraint type
type Number interface {
	int64 | float64
}

func main() {
	xInt := 1
	yInt := 5
	z := sum.Add(xInt, yInt)
	fmt.Println(z)

	xFloat := 2.5
	yFloat := 5.5

	zFloat := sum.Add(xFloat, yFloat)
	fmt.Println(zFloat)

	greet := "hello "
	who := "lizzy mcalpine"

	greetings := sum.Add(greet, who)
	fmt.Println(greetings)

	var a, b = 5, 4

	minimum := min.Min(a, b)
	fmt.Println(minimum)

	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n", sumInts(ints), sumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n", sumIntOrFloats(ints), sumIntOrFloats(floats))

	fmt.Println(ItsTheSame(1, 3))

	fmt.Println("========NON-GENERIC STRUCT=========")
	structSumInts := operation.New(5, 10)
	structSumInts.Sum()
	fmt.Printf("Sum int struct non-generic: %v\n", structSumInts.Result)

	structSumFloats := operation.NewFloat(5.5, 10.5)
	structSumFloats.Sum()
	fmt.Printf("Sum floats struct non-generic: %v\n", structSumFloats.Result)

	fmt.Println("========GENERIC STRUCT=========")

	calcsIntsOrFloats := operation.NewCalculatorIntsOrFloats(3.2, 5.10)
	calcsIntsOrFloats.Sum()

	fmt.Printf("Sum floats struct generic: %v\n", calcsIntsOrFloats.Result)
}

func sumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}

	return s

}

func sumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}

	return s
}

func sumIntOrFloats[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}

func ItsTheSame[T comparable](a, b T) bool {

	return a == b
}
