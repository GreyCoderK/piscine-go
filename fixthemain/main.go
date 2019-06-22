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

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...\n")
	ptrDoor.State = CLOSE
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("Door is open ?\n")
	return ptrDoor.State == OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("Door is close ?\n")
	return ptrDoor.State == CLOSE
}

func OpenDoor(ptrDoor *Door){
	PrintStr("Door Opening...\n")
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
