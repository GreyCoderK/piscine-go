package piscine

func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	for i:=0;i < len(array) - 1; i++ {
		for j:= i+1; j < len(array); j++{
			if f(array[i],array[j]) == 1 {
				temp := array[i]
				modify(array, i, array[j])
				modify(array, j, temp)
			}
		}
	}	
}

func modify(array []string, ind int, val string){
	array[ind] = val
}
