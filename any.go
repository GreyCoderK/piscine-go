package piscine

func Any(f func(string) bool, arr []string) bool {
	for _,res:= range arr {
		if f(res) {
			return true
		}
	}
	return false
}
