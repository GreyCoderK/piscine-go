package piscine

import "strconv"

func BasicAtoi(s string) int {
	i, err := strconv.Atoi(s)
	
	if err == nil {
		return i
	}else{
		return 0
	}
}
