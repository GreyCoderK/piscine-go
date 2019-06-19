package main

import (
    "os"
    "fmt"
    "io/ioutil"
)

func readFile(s string) {
    file, err := os.Open(s)
    if err != nil {
        fmt.Println(err.Error())
    }

    data, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err.Error())
    }

    fmt.Printf("%s\n\n", data)
    file.Close()
}

func main(){
	args := os.Args[1:]
	if 0 == len(args) {
		fmt.Println(os.Args[0])
	}else{
		for _,res:= range os.Args[1:] {
			readFile(res)
		}
	}
}
