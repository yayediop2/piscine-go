package piscine

func IsAlpha(s string) bool {
	bs := []byte(s)
	for _, value := range bs {
		if (value < 48 || value > 57) && (value < 65 || value > 90) && (value < 97 || value > 122) {
			return false
		}
	}
	return true
}