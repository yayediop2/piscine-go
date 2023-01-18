package main

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var tab []int
	if n == 0 {
		z01.PrintRune('0')
	}
	for n != 0 {
		mod := n % 10
		tab = append(tab, mod)
		n = n / 10
	}
	for i := 0; i <= len(tab)-1; i++ {
		for j := 0; j <= len(tab)-1; j++ {
			if tab[i] < tab[j] {
				tab[i], tab[j] = tab[j], tab[i]
			}
		}
	}
	for index := range tab {
		z01.PrintRune(rune(tab[index] + '0'))
	}
}

func main() {
	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)
}
