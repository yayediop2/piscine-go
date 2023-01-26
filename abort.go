package piscine

func Abort(a, b, c, d, e int) int {
	v := []int{a, b, c, d, e}
	var z int
	for i := 0; i < len(v); i++ {
		for j := i + 1; j < len(v); j++ {
			if v[i] > v[j] {
				v[i], v[j] = v[j], v[i]
				z = v[2]
			}
		}
	}
	return z
}

/*
func main() {
	middle := Abort(-2, -3, -889654, -674875, -7)
	fmt.Println(middle)
}
*/
