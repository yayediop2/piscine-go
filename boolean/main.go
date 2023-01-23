package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func main() {
	arg := os.Args
	if len(arg)%2 == 1 {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
