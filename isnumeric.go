package piscine

func IsNumeric(s string) bool {
	bs := []byte(s)
	for _, value := range bs {
		if value < 48 || value > 57 {
			return false
		}
	}
	return true
}
