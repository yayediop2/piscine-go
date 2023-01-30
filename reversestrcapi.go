package main

import (
	"os"

	"github.com/01-edu/z01"
)

// if sr[i]=' ' && sr[i-1]>='a' &&  sr[i-1]>='z'

func ReverseStrCapi() {
	arg := os.Args[1:]
	if len(arg) >= 1 {
		for i := 0; i < len(arg); i++ {
			sr := []rune(arg[i])
			for j := 0; j < len(sr); j++ {
				if sr[j] >= 'A' && sr[j] <= 'Z' {
					sr[j] = sr[j] + ' '
				}
				if sr[j] == ' ' && sr[j-1] >= 'a' && sr[j-1] <= 'z' {
					sr[j-1] = sr[j-1] - ' '
				}
			}
			l := sr[len(sr)-1]
			if l >= 'a' && l <= 'z' {
				sr[len(sr)-1] = sr[len(sr)-1] - ' '
			}
			PrintStr(sr)
		}
	}
}

func PrintStr(a []rune) {
	for _, v := range a {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

func main() {
	ReverseStrCapi()
}
