package piscine

import "unicode"

func Isupper(str string) bool {
	res := []rune(str)
	return unicode.IsUpper(res)
}
