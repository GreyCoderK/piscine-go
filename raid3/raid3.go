package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
	"os/exec"
)

func main(){
	info, err := os.Stdin.Stat()
    	if err != nil {
        	fmt.Println(err.Error())
    	}
	
	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		fmt.Println()
		return
	}
	
	reader := bufio.NewReader(os.Stdin)
	var output []rune
	
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}
	
	row, col:= getDimension(string(output))
	fmt.Println(IsRaid1a(row, col))
}

func IsRaid1a (row, col int) string {
	cmd, _ := exec.Command("raid1a", string(row), string(col)).Output()
	return string(cmd)
}


func getDimension(s string) (int, int) {
	col, row:= 0, 0
	for _, res := range s {
		if row == 0 && res != '\n' {
			col++
		}
		if res == '\n' {
			row++
		}
	}
	return row, col
}
