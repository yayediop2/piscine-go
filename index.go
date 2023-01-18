package piscine

func Index(s string, toFind string) int {
	a := len(s)
	b := len(toFind)
	if b > a {
		return -1
	}
	for i := 0; i <= a-b; i++ {
		if toFind == s[i:i+b] {
			return i
		}
	}

	return -1
}
