package piscine

import "strings"

func Capitalize(s string) string {
	return strings.Replace(strings.Title(ToLower(s)), s[strings.Index("_")+1], ToUpper(s[strings.Index("_")+1]))
}
