package piscine

import "unicode"

func IsAlpha(str string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return true
		}
	}
	return false
}
