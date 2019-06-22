package piscine

func Compact(ptr *[]string, length int) int {
	compte := 0
	for i:=0; i < length; i++{
		if (*ptr)[i] != " " {
			compte++
		}
	}
	return compte
}
