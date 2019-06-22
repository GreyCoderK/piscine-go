package main

import (
	"fmt"
	piscine ".."
)

func main() {
	var donnie piscine.Pilot
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = piscine.AIRCRAFT1

	fmt.Println(donnie)
}
