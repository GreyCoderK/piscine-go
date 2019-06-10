package piscine

import "fmt"

func Raid1a(x,y int) {
	i, j := 0, 0
	for i <= y {
		for j <= x {
			if i == 0 || i == y - 1 {
				if j == 0 || j == x - 1 {
					fmt.Print("o")
				}else{
					fmt.Print("-")
				}
			}else{
				if j == 0 || j == x - 1 {
					fmt.Print("|")
				}else{
					fmt.Print(" ")
				}
			}
			if j == x - 1 {
				fmt.Print("\n")		
			}
			j++
		}
		i++
	}
}
