package piscine

import "unicode"

func IsPrintable(str string) bool {
	res := []rune(str)
	return unicode.IsPrint(res)
}
