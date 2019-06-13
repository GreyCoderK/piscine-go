package piscine

import "unicode"

func IsPrintable(str string) bool {
	return unicode.IsPrint(str)
}
