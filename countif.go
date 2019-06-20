package piscine

func CountIf(f func(string) bool, arr []string) int {
        compte := 0
	for _,res:= range arr {
                if f(res) {
                        compte += 1
                }
        }
        return compte
}
