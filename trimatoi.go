package piscine

func TrimAtoi(s string) int {
	p := []rune(s)
	k := 1
	var x int = 0
	for i := 0; i < len(s); i++ {
		if p[i] == '-' && x == 0 {
			k = -1
		}
		if p[i] >= '0' && p[i] <= '9' {
			x = x*10 + int(p[i]-'0')
		}

	}
	return k * x
}

/* }
for _, v := range bs {
	if v >= '0' && v <= '9' {
		e = e*10 + int(v+'0')
	}
	if v == '-' && e == 0 {
		signe = -1
	}
} */

/* main

func main() {
	fmt.Println(TrimAtoi("12345"))
	fmt.Println(TrimAtoi("str123ing45"))
	fmt.Println(TrimAtoi("012 345"))
	fmt.Println(TrimAtoi("Hello World!"))
	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
} */
