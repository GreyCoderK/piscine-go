package piscine

import (
	"fmt"
)

func Raid1C(x, y int){
	if x == 1 && y == 1 {
		fmt.Print("A\n")
	}else{
		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				if i == 0 || i == y-1 {
					if i == 0 && j == 0 || i == 0 && j == x-1 {
						fmt.Print("A")
					}else if i == y -1 && j == 0 || i == y -1 && j == x-1 {
						fmt.Print("C")
					} else {
						fmt.Print("B")
					}
				} else {
					if j == 0 || j == x-1 {
						fmt.Print("B")
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
}