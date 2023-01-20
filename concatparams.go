package piscine

func ConcatParams(args []string) string {
	var t string
	for i := range args {
		if i == len(args)-1 {
			t = t + args[i]
		} else {
			t = t + args[i] + "\n"
		}
	}
	return t
}
