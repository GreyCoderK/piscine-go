package piscine

func FindNextPrime(nb int) int {
	if piscine.IsPrime(nb) {
		return nb
	} else {
		n := nb + 1
		for piscine.IsPrime(n) == false {
			n++
		}
		return n
	}
}
