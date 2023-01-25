package piscine

func Any(f func(string) bool, a []string) bool {
	for _, v := range a {
		if f(v) {
			return true
		}
	}
	return false
}

/* func main() {
	a1 := []string{"Hello", "how", "are", "you"}
	a2 := []string{"This", "is", "4", "you"}

	result1 := Any(IsNumeric, a1)
	result2 := Any(IsNumeric, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
*/
/* func IsNumeric(s string) bool {
	bs := []rune(s)
	for _, value := range bs {
		if value < '0' || value > '9' {
			return false
		}
	}
	return true
} */
