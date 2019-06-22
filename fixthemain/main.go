package main 

import "github.com/01-edu/z01"

const (
	CLOSE = 0
	OPEN = 1
)

type Door struct {
	State int
}

func PrintStr(str string) {
	arrayRune := []rune(str)
	for _, s := range arrayRune {
		z01.PrintRune(s)
	}
}

func CloseDoor(ptrDoor *Door) bool{
	PrintStr("Door Closing...\n")
	ptrDoor.State = CLOSE
	return true
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?\n")
	return ptrDoor.State == OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?\n")
	return ptrDoor.State == CLOSE
}

func OpenDoor(ptrDoor *Door){
	ptrDoor.State = OPEN
}

func main() {
	var door Door

	OpenDoor(&door)
	if IsDoorClose(&door) {
		OpenDoor(&door)
	}
	if IsDoorOpen(&door) {
		CloseDoor(&door)
	}
	if door.State == OPEN {
		CloseDoor(&door)
	}
}
