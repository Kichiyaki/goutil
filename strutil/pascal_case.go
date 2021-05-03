package strutil

import (
	"unicode"
)

func PascalCase(s string, separators ...rune) string {
	n := ""
	upperNext := true

loop:
	for _, c := range s {
		for _, sep := range separators {
			if c == sep {
				upperNext = true
				continue loop
			}
		}
		if upperNext {
			if unicode.IsLower(c) {
				c = unicode.ToUpper(c)
			}
			upperNext = false
		}
		n += string(c)
	}
	return n
}
