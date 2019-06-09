package piscine

import "strconv"

func BasicAtoi2(s string) int {
	i, err := strconv.Atoi(s)
	
	if err == nil {
		return i
	}else{
		return 0
	}
}
