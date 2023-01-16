package piscine

func IterativeFactorial(nb int) int {
	var r int = 1

	if nb == 0 {
		return 1
	}
	if nb < 0 || nb > 20 {
		return 0
	}
	if nb > 0 {
		for i := 1; i <= nb; i++ {
			r = r * i
		}
	}
	return r
}
