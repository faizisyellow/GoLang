package main

func createTransform(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}
