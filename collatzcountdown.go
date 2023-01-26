package piscine

func CollatzCountdown(start int) int {
	c := 0
	if start <= 0 {
		return -1
	} else {
		for i := 0; i <= start; i++ {
			if start%2 == 0 {
				start = start / 2
				c++
			}
			if start%2 == 1 {
				c++
				start = (start * 3) + 1
				c++
			}
		}
	}
	return c
}

/* func main() {
	steps := CollatzCountDown(12)
	fmt.Println(steps)
}
*/
