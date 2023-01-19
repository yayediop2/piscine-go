package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := []rune(os.Args[0])
	for i := 2; i < len(a); i++ { // les deux premiers elements sont le point et le slash
		z01.PrintRune(a[i])
	}
	z01.PrintRune('\n')
}
