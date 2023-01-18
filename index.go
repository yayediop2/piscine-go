package piscine

func Index(s string, toFind string) int {
	a := len(toFind)
	for i := range s {
		if s[i] == toFind[0] {
			if len(s) > i+a {
				if toFind == s[i:i+a] {
					return i
				}
			}
		}
	}
	return -1
}
