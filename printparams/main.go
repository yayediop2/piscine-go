package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for i := 1; i < len(arg); i++ {
		for _, v := range arg[i] {
			z01.PrintRune(rune(v))
		}
		z01.PrintRune('\n')
	}
}
