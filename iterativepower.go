package piscine

func IterativePower(nb, power int) int {
	if 1 == power {
		return nb
	} else if 0 == power {
		return 1
	} else {
		res := 1
		for i := 0; i < power; i++ {
			res *= nb
		}
		return res
	}
}
