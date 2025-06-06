package main

import "github.com/01-edu/z01"

const (
	OPEN  = true
	CLOSE = false
)

type Door struct {
	state bool //
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
} // neat

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...\n")

	ptrDoor.state = CLOSE //
	return true
}

func OpenDoor(ptrDoor *Door) bool {
	PrintStr("Door Opening...\n")
	ptrDoor.state = OPEN //
	return true
}

func IsDoorOpen(ptrDoor *Door) bool { //
	PrintStr("is the Door opened ?\n")
	return ptrDoor.state == OPEN //
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?\n")
	return ptrDoor.state == CLOSE
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}
