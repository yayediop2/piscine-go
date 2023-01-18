package piscine

func IsLower(s string) bool {
	bs := []byte(s)
	for _, letter := range bs {
		if letter < 97 || letter > 122 {
			return false
		}
	}
	return true
}
