package main

import (
	"github.com/01-edu/z01"
)

type DoorState int

const (
	OPEN  DoorState = 0
	CLOSE DoorState = 1
)

type Door struct {
	state DoorState
}

func OpenDoor(ptrDoor *Door) {
	message := "Door Opening..."
	printMessage(message)
	ptrDoor.state = OPEN
}

func CloseDoor(ptrDoor *Door) {
	message := "\nDoor Closing..."
	printMessage(message)
	ptrDoor.state = CLOSE
}

func IsDoorOpen(door *Door) bool {
	return door.state == OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
	return ptrDoor.state == CLOSE
}

func printMessage(message string) {
	for _, char := range message {
		z01.PrintRune(char)
	}
}

func main() {
	door := &Door{}

	OpenDoor(door)
	printMessage("\nis the Door closed ?")
	if IsDoorClose(door) {
		printMessage("\nThe Door is closed.")
	}
	printMessage("\nis the Door opened ?")
	CloseDoor(door)
	printMessage("\n")
}
