package main

func fibonacci() func() int {
	a, b := 0, 1

	closure := func() int {
		result := &a

		a, b = b, a+b
		return *result
	}
	return closure
}
