package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	neg := 1
	if n < 0 {
		z01.PrintRune('-')
		neg = -1
	}

	if n != 0 {
		nouveauN := n / 10 * neg // pour récupérer le résultat de la division
		if nouveauN != 0 {       // ça s'arrete lorsqu'on arrive à 0
			PrintNbr(nouveauN) //
		}

	}
	result := n % 10 * neg
	z01.PrintRune(rune(result) + '0')
}
