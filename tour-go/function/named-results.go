package main

func split(sum int) (x, y int, message string) {
	x = sum * 4 / 9
	y = sum - x

	message = "hello"

	// 'naked' return. (it returns the named return values.)
	return
}
