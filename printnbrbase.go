package piscine

import(
	"math"
	"fmt"
	"strings"
)

func Reverse(s string) string {
    var reverse string
    for i := len(s)-1; i >= 0; i-- {
        reverse += string(s[i])
    }
    return reverse 
}

func PrintNbrBase(nbr int32, str string)(){
	indx := 0
	
	for _,res:= range str {
		if string(res) == "-" || string(res) == "+" || strings.Count(str, string(res)) > 1 {
			indx = 1
			break
		}
	}
	if math.MaxInt32 < nbr || math.MinInt32 > nbr{
		fmt.Println(nbr)
	}else if indx == 1 || len(str) <= 2{
		fmt.Println("NV")	
	}else{
		if nbr < 0 {
			fmt.Print("-")
			nbr *= -1	
		}
		i:=0
		nan:=""
		for nbr >= int32(len(str)) {
			if nbr >= int32(len(str)) {
				nan +=string(str[nbr % int32(len(str))])
				nbr = nbr/int32(len(str))
				i++
			}
			
		}
		
		nan +=string(str[nbr])
		fmt.Println(Reverse(nan))
	}
}
