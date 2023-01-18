package piscine

func IsUpper(s string) bool {
	bs := []byte(s)
	for _, letter := range bs {
		if letter < 65 || letter > 90 {
			return false
		}
	}
	return true
}