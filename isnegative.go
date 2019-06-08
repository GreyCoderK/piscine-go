package piscine

import "fmt"

func IsNegative(nb int) {
	res := nb > 0
	if res {
		fmt.Println('T')
	}else{
		fmt.Println('F')
	}
}
