package strutil

import (
	"unicode"
)

func PascalCase(s string, sep rune) string {
	n := ""
	upperNext := true
	for _, c := range s {
		if c == sep {
			upperNext = true
			continue
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
