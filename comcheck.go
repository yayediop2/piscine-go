package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	for _, v := range arg {
		if v == "galaxy" || v == "01" || v == "galaxy 01" {
			com := "Alert!!!"
			for _, val := range com {
				z01.PrintRune(val)
			}
			z01.PrintRune('\n')
		}
	}
}
