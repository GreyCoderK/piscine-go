package piscine

func Max(arr []int) int {
	max := arr[0]
	for i:=1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}
