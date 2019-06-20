package main

import (
    "os"
    "fmt"
    "io/ioutil"
    "strconv"
)

func main(){
	args := os.Args[1:]
        if 0 == len(args) {
                fmt.Println(os.Args[0])
        }else{
		num, err := strconv.Atoi(os.Args[1])
	
		if err != nil {
			fmt.Println(err.Error())
		}else{
			if 0 == len(os.Args[2:]) {
				fmt.Println("File name missing")
			}else{
				for _,res:= range os.Args[2:] {
                                	ztail(res, num)
                        	}	
			}
		}
        }
}

func ztail(s string, numBytes int) {
    file, err := os.Open(s)
    if err != nil {
        fmt.Println(err.Error())
    }

    data, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err.Error())
    }

    fmt.Printf("%s\n", data[len(data)-numBytes:])
}
