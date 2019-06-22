package main

import(
	"os"
	"fmt"
)

func main(){
	res := ""
	args = os.Args[1:]
	check := []string{"01", "galaxy", "galaxy 01"}
	for _, res:= range args {
		for _,item := range check {
			if res == item {
				res += "Alert!!!"
				break
			}
		}
	}
	res +="\n"
	fmt.Print(res)
}
