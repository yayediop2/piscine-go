package main

import "fmt"

func ponctuation(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] == '.' || s[i] == ',' || s[i] == '!' || s[i] == '?' || s[i] == ':' || s[i] == ';' {
			if s[i+1] == ' ' && s[i] != s[len(s)-2] {
				i++
			}
		}
	}
	return s
}

func main() {
	a := ponctuation("jtsy. guu:;,l")
	fmt.Print(a)
}
