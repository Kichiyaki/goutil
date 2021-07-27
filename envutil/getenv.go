package envutil

import (
	"os"
	"strconv"
)

func GetenvInt(key string) int {
	str := GetenvString(key)
	if str == "" {
		return 0
	}
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

func GetenvBool(key string) bool {
	b, _ := strconv.ParseBool(GetenvString(key))
	return b
}

func GetenvString(key string) string {
	return os.Getenv(key)
}
