package words

import (
	"unicode"
)

type Str string

func (s *Str) UpperOne() {
	var newStr Str

	for i, v := range *s {
		if i == 0 {
			v = unicode.ToUpper(v)
			newStr += Str(v)
		} else {
			newStr += Str(v)
		}
	}

	*s = newStr
}
