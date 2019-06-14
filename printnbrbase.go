package piscine

import (
	"fmt"
	"strings"
)

func PrintNbrBase(nbr int, str string){
	indx := 0
	
	for _,res:= range str {
		if string(res) == "-" || string(res) == "+" || strings.Count(str, string(res)) > 1{
			indx = 1
			break
		}
	}
	
	if indx == 1 {
		fmt.Print("NV")	
	}else{
		if nbr < 0 {
			fmt.Print("-")
			nbr *= -1	
		}
		i:=0
		res := ""
		for nbr > len(str) {
			if nbr > len(str) {
				res += string(str[nbr % len(str)])
				nbr = int(nbr/len(str))
				i++
			}
			
		}
		
		res += string(str[nbr])
		fmt.Print(reverse(res))
	}
}

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
