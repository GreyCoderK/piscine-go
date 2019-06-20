package piscine

func ForEach(f func(int), arr []int) {
	for _,res:= range arr{
		f(res)
	}
}
