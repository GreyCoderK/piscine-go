package piscine

import "sort"

func Abort(a, b, c, d, e int) int {
	tab := []int{a, b, c, d, e}
	sort.IntSlice(tab)
	return tab[2]
}

