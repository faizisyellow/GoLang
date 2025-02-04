package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	var steps int

	if n <= 0 {
		return 0, errors.New("only positive numbers are allowed")
	}
	if n == 1 {
		return 0, nil
	}

	for i := n; i > 0; i-- {
		if n%2 == 0 {
			n = n / 2
			steps++
		} else {
			n = n*3 + 1
			steps++
		}

		// if n has reached 1, out from loop
		if n == 1 {
			break
		}
	}

	return steps, nil
}

// SOLUTION 2.

/*
 func CollatzConjecture(n int) (int, error) {
	var steps int

	if n <= 0 {
		return 0, errors.New("only positive numbers are allowed")
	}
	if n == 1 {
		return 0, nil
	}

	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}

		steps++
	}

	return steps, nil
 }
*/
