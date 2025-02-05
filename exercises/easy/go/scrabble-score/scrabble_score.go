package scrabble

import "unicode"

func Score(word string) int {
	points := []int{1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3, 1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10}
	var score int

	for _, v := range word {
		if unicode.IsUpper(v) {
			score += points[v-'A']
		} else {
			score += points[v-'a']
		}
	}

	return score
}
