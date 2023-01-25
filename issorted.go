package piscine

/* func isS(a, b int) int {
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
} */

func IsSorted(f func(a, b int) int, a []int) bool {
	if a[0] < len(a)-1 {
		for i := 0; i < len(a); i++ {
			for j := i + 1; j < len(a); j++ {
				if f(a[i], a[j]) > 0 {
					return false
				}
			}
		}
	} else if a[0] > len(a)-1 {
		for i := 0; i < len(a); i++ {
			for j := i + 1; j < len(a); j++ {
				if f(a[i], a[j]) < 0 {
					return false
				}
			}
		}
	}
	return true
}

/*
func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}

	result1 := IsSorted(isS, a1)
	result2 := IsSorted(isS, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
*/
