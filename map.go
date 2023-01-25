package piscine

func Map(f func(int) bool, a []int) []bool {
	rslice := make([]bool, len(a))
	for i, v := range a {
		rslice[i] = f(v)
	}
	return rslice
}
