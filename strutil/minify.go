package strutil

import "strings"

func Minify(val string) string {
	return strings.Join(strings.Fields(val), " ")
}
