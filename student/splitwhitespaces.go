package piscine

import "strings"

func SplitWhiteSpaces(str string) []string {
	tab := strings.Split(str, " ")
	return tab
}
