package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	if len(arg) == 2 {
		for _, letter := range arg[1] {
			if letter >= 'a' && letter <= 'z' {
				for i := 'a'; i <= letter; i++ {
					z01.PrintRune(letter)
				}
			}
			if letter >= 'A' && letter <= 'Z' {
				for i := 'A'; i <= letter; i++ {
					z01.PrintRune(letter)
				}
			}
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}
