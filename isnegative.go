package piscine

import "fmt"

func IsNegative(nb int) {
	res := nb > 0
	if res {
		fmt.Print("%c\n","T")
	}else{
		fmt.Print("%c\n","F")
	}
}
