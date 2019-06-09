package piscine

import (
    "fmt"
    "strconv"
)

func PrintComb() {
for i:=0; i<10; i++{ 
    	for j:=i+1; j<10; j++{
         	for k:=j+1; k<10; k++{
			if i == 0 && j == 1 && k == 2 {
				fmt.Print("")
			}else{
				fmt.Print(", ")
			}
			ch := strconv.Itoa(i)
			ch += strconv.Itoa(j)
			ch += strconv.Itoa(k)
             		fmt.Print(ch)
        	}
    	}
    }
}
