package min

import "cmp"

// Ordered type (type that support comparison operation)
func Min[T cmp.Ordered](x, y T) T {
	if x < y {
		return x
	}

	return y
}
