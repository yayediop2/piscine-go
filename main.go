package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func printprogramname() {
	a := os.Args[0]
	for _, v := range a {
		z01.PrintRune(rune(v))
	}
}

/*
func main() {
	printprogramname()
}
*/
