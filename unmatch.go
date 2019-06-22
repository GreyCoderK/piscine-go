package piscine

func Unmatch(arr []int) int {
	for _,res:= range arr {
		fois:= 0
		for _,el:= range arr {
			if el == res {
				fois++
			}
		}
		if fois == 1 || fois % 2 == 1 {
			return res
		}
	}
	return -1
}
