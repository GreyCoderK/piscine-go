package piscine

import "unicode"

func IsLower(str string) bool {
	res := []rune(str)
	return unicode.IsLower(res)
}
