package main

import (
	"os"
	"fmt"
	"strconv"
)

func validateOperator(test string) bool {
	op := []string{"+", "-", "*", "/", "%"}
	for _, res:= range op {
		if res == test {
			return true
		}
	}
	return false
}

func main(){
	args := os.Args[1:]
	if len(args) > 3 || len(args) < 3 {
		fmt.Print()
	}else{
		if validateOperator(args[1]) == false {
			fmt.Println(0)
		}else{
			premier, _ := strconv.Atoi(args[0])
			second, _ := strconv.Atoi(args[2])
			
			if args[1] == "%" && second == 0 {
				fmt.Print("No Modulo by 0\n")
			}else if args[1] == "/" && second == 0 {
				fmt.Print("No division by 0\n")
			}else if args[1] == "+" {
				fmt.Println(premier+second)
			}else if args[1] == "-" {
				fmt.Println(premier-second)
			}else if args[1] == "*" {
				fmt.Println(premier * second)
			}else if args[1] == "/" {
				fmt.Println(premier/second)
			}else{
				fmt.Println(premier%second)
			}

		}
	}

}
