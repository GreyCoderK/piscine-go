package piscine

import (
	"strings"
)

func AtoiBase(s string, base string) int {
	for _,res:= range str {
		if string(res) == "-" || string(res) == "+" || strings.Count(str, string(res)) > 1 {
			indx = 1
			break
		}
	}
	if indx == 1 || len(str) < 2{
		return 0	
	}else{
		fin := 0
		for i,res:= range s {
			ind = stings.Index(base,res)
			fin += RecursivePower(i,index + 1)
		}
		return fin
	}
}
