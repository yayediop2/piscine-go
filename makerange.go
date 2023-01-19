package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	z := max - min
	a := make([]int, z)

	for i := 0; i < z; i++ {
		a[i] = min + i
	}
	return a
}
