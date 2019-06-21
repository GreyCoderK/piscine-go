package main

import (
    "os"
    "fmt"
    "io/ioutil"
)

func main(){
	args := os.Args[1:]
	if 0 == len(args) {
		fmt.Print()
	}else{
		for _,s:= range os.Args[1:] {
			file, err := os.Open(s)
    			if err != nil {
    			    fmt.Println(err.Error())
    			}else{
        			data, err := ioutil.ReadAll(file)
        			if err != nil {
                 			fmt.Println(err.Error())
					break outer
         			}else{
         				fmt.Printf("%s", data)
				}
    			}
			file.Close()
		}
	}
}
