package piscine

import "strings"

func Capitalize(s string) string {
	res := strings.Title(ToLower(s))
	indx := strings.Index("_")
	if indx == -1 {
		return res
	} else {
		res = strings.Replace(res, res[indx+1], ToUpper(res[indx+1]), indx)
		return res
	}
}
