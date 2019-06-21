package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	for i:= 0; i < len(tab) - 1; i++{
		if f(tab[i],tab[i+1]) == -1 {
			return false	
		}
	}
	return true
}
