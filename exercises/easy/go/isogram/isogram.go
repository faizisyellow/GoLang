package isogram

import (
	"slices"
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	var chars []rune

	for _, v := range word {
		// space should be passed, and only letter got checked.
		if slices.Contains(chars, v) && !unicode.IsSpace(v) && unicode.IsLetter(v) {
			return false
		}
		chars = append(chars, v)
	}
	return true
}
