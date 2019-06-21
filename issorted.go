package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	ascendant, descendant:= true,true
	for i:= 0; i < len(tab) - 1; i++{
		if f(tab[i],tab[i+1]) == -1 {
			ascendant = false	
		}
	}

	for i:=0; i < len(tab); i++ {
		if f(tab[i],tab[i+1]) == 1 {
                        descendant = false
                }
	}
	return acendant || descendant
}
