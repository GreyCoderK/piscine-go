package piscine

import "fmt"

func Raid1b(x, y int){
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 || i == y-1 {
				if j == 0 && i==0 || j == x-1 && i == y-1 {
					fmt.Print("/")
				}else if j == x-1 && i==0 || j == 0 && i == y-1 {
					fmt.Print("\\")
				} else {
					fmt.Print("-")
				}
			} else {
				if j == 0 || j == x-1 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
			if j == x-1 {
				fmt.Print("\n")
			}
		}
	}
}