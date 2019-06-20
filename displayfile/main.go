package main

import (
	"os"
	"fmt"
	"strings"
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
			fmt.Printf("%s", strings.Replace(err.Error(), "\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\n", "\n", -1))
		}
		arr := make([]byte,14)
		file.Read(arr)
		fmt.Println(string(arr))
		file.Close()
	}
}
