package piscine

func IterativePower(nb int, power int) int {
	var r int = 1
	if power < 0 {
		return 0
	}

	for i := 1; i <= power; i++ {
		r = r * nb
	}
	return r
}
