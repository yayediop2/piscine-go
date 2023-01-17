package piscine

func FirstRune(s string) rune {
	var i byte

	for range s {
		i = s[0]
	}
	return rune(i)
}
