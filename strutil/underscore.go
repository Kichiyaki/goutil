package strutil

import (
	"unicode"
)

func Underscore(s string) string {
	r := ""
	for i, c := range s {
		if unicode.IsUpper(c) {
			if i > 0 && i+1 < len(s) && (unicode.IsLower(rune(s[i-1])) || unicode.IsLower(rune(s[i+1]))) {
				r += "_"
			}
			c = unicode.ToLower(c)
		}
		r += string(c)
	}
	return r
}
