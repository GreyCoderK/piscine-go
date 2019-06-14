package piscine

import (
	"fmt"
	"strings"
)

func PrintNbrBase(nbr int, str string)(){
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
		for nbr > len(str) {
			if nbr > len(str) {
				defer fmt.Print(string(str[nbr % len(str)]))
				nbr = nbr/len(str)
				i++
			}
			
		}
		
		defer fmt.Print(string(str[nbr]))
		
	}
}
