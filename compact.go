package piscine

func Compact(ptr *[]string, length int) int {
	cmp := 0
	for_, res := range ptr {
		if *(res) != " " {
			cmp++
		}
	}
	return cmp
}
