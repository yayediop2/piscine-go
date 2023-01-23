package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func Print_int(a int) {
	if a == 0 {
		z01.PrintRune('0')
	}
	if a != 0 {
		na := a / 10
		if na != 0 {
			Print_int(na)
		}
	}
	r := a % 10
	z01.PrintRune(rune(r) + '0')
}

func main() {
	points := &point{}
	setPoint(points)
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	Print_int(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	Print_int(points.y)
	z01.PrintRune('\n')
}
