package main

import (
	"os"
	"fmt"
)

func main(){
	args := os.Args[1:]
	if 0 == len(args) {
		fmt.Println("File name missing")
	}else if len(args) >= 2 {
		fmt.Println("Too many arguments")
	}else{
		file, err:= os.Open(string(args[0]))
		
		if err != nil {
			fmt.Print(err.Error())
		}
		arr := make([]byte,14)
		file.Read(arr)
		fmt.Println(string(arr))
		file.Close()
	}
}
