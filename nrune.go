package piscine

func NRune(s string, n int) rune {
	m := []rune(s)
	for index := range m {
		if n == index+1 {
			return m[index]
		}
	}
	return 0
}
