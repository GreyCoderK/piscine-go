package piscine

import (
    "bytes"
    "fmt"
    "strconv"
)

func PrintComb() {
    var buffer bytes.Buffer

    for i := 12; i < 1000; i++ {
	a := int(i/100)
	b := int(i/10)
	c := int(i%10)
	if(a < b && b < c){
		if(i>12) {
			buffer.WriteString(", ")
		}
        	buffer.WriteString(strconv.Itoa(a))
		buffer.WriteString(strconv.Itoa(b))
		buffer.WriteString(strconv.Itoa(c))
	}
    }

    fmt.Println(buffer.String())
}
