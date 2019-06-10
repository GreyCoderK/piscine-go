package piscine

import "fmt"

func Raid1a(x, y int) {
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 || i == y-1 {
				if j == 0 || j == x-1 {
					fmt.Print("o")
				} else {
					fmt.Print("-")
				}
			} else {
				if j == 0 || j == x-1 {
					fmt.Print("|")
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
