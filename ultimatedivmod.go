package piscine

func DivMod(div *int, mod *int) {
	x := *div
	*div = *div / *mod
	*mod = x % *mod
}
