package strutil

import "strings"

func Minify(val string, sep string) string {
	return strings.Join(strings.Fields(val), sep)
}
