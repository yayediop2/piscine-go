package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func Print_int(a int) {
	nouveauN := a / 10
	if nouveauN != 0 {
		Print_int(nouveauN)
	}
	result := a % 10
	z01.PrintRune(rune(result) + '0')
}

func main() {
	points := point{}
	points.x = 42
	points.y = 21
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
