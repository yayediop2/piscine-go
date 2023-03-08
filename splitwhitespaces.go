package main

import "fmt"

func SplitWhiteSpaces(s string) []string {
	tab := []string{}
	h := ""
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			h = h + string(s[i])
		}

		if s[i] == ' ' || s[i] == '\t' || s[i] == '\n' || s[i] == s[len(s)-1] {
			if h != "" {
				tab = append(tab, h)
				h = ""
			}
		}
	}

	return tab
}

func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello    how   are you?"))
}
