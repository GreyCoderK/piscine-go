package piscine

import "strings"

func Capitalize(s string) string {
	return strings.Title(ToUpper(s))
}
