package hamming

import (
	"errors"
	"strings"
)

func Distance(a, b string) (int, error) {
	var distance int
	x := strings.Split(a, "")
	y := strings.Split(b, "")

	if len(a) != len(b) {
		return 0, errors.New("invalid")
	}

	for i, v := range x {
		if v != y[i] {
			distance++
		} else {
			continue
		}

	}

	return distance, nil

}

/*
	Solution 2.

	func Distance(a, b string) (int, error) {
		var distance int

		if len(a) != len(b) {
			return 0, errors.New("invalid")
		}

		for i := range a {
			if a[i]!=b[i] {
				distance++
				}
			}

			return distance, nil

		}

*/
