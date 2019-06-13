package piscine

import "unicode"

func Isupper(str string) bool {
	for _, res := range str {
		if !unicode.IsUpper(res) {
			return false
		}
	}
	return true
}
