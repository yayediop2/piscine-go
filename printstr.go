package piscine

func PrintStr(s string) {
	for i := 0; i < len(s); i++ {
		PrintRune(rune(s[i]))
	}
}
