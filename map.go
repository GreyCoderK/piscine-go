package piscine

func Map(f func(int) bool, arr []int) []bool {
	
	fin:= make([]bool, len(arr))
	
	for i, res:= range arr {
		fin[i] = f(res)
	}
	
	return fin
}
