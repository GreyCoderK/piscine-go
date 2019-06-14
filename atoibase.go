package piscine

import (
	"strings"
)

func AtoiBase(s string, str string) int {
	indx:=0
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
			ind := strings.Index(base,res)
			fin += RecursivePower(i,ind + 1)
		}
		return fin
	}
}
