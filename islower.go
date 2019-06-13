package piscine

import "unicode"

func IsLower(str string) bool {
	res := []rune(str)
	for _, res := range str {
		if !unicode.IsLower(res) {
			return false
		}
	}
	return true
}
