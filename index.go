package piscine

func Index(s string, toFind string) int {
	for index := range s {
		if toFind[0] == s[index] {
			return index
		}
	}
	return -1
}
