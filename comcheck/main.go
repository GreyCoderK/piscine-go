package main

import(
	"os"
	"fmt"
)

func main(){
	result := ""
	args := os.Args[1:]
	check := []string{"01", "galaxy", "galaxy 01"}
	for _, res:= range args {
		for _,item := range check {
			if res == item {
				result += "Alert!!!\n"
				break
			}
		}
	}
	fmt.Print(result)
}
