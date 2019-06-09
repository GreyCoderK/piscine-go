package piscine

import (
    "fmt"
    "strconv"
)

func PrintComb2() {
	for i:=0; i<99; i++{ 
    		for j:=i+1; j<100; j++{
			if i == 0 && j == 1 {
				fmt.Print("")
			}else{
				fmt.Print(", ")
			}
			if i < 10 {
				ch := strconv.Itoa(0)
				ch += strconv.Itoa(i)
				fmt.Print(ch," ")
             		}else{
				fmt.Print(i," ")
			}
			
			if j < 10 {
                                ch := strconv.Itoa(0)
                                ch += strconv.Itoa(j)
                                fmt.Print(ch)
                        }else{ 
                                fmt.Print(j)
                        }
			if i==98 && j== 99{
				print("\n")
			}
			
        	}
    	}
}
