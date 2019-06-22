package piscine

import "sort"

func Abort(a, b, c, d, e int) int {
	tab := []int{a, b, c, d, e}
	res := sort.IntSlice(tab)
	return res[2]
}

