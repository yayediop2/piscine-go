package piscine

func LastRune(s string) rune {
	m := []rune(s)
	i := 0
	for index := range s {
		i = index
	}
	return m[i]
}
