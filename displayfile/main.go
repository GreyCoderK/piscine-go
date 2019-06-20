package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

func main(){
	args := os.Args[1:]
	if 0 == len(args) {
		fmt.Println("File name missing")
	}else if len(args) >= 2 {
		fmt.Println("Too many arguments")
	}else{
		file, err:= ioutil.ReadFile(args[0])
		
		if err != nil {
			fmt.Print(err)
		}
		return string(file)
	}
}
