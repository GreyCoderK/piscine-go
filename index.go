package piscine

func Index(s string, toFind string) int {
	cmp := 0
	res := []rune(s)
	for i := 0; i < len(res); i++ {
		if toFind == string(res[i]) {
			cmp++
		}
	}
	if cmp == 0 {
		return -1
	} else {
		return cmp
	}
}
