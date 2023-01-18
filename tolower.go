package main

import "fmt"

func ToLower(s string) string {
	bs := []rune(s)
	for i, value := range bs {
		if value >= 65 && value <= 90 {
			bs[i] += 32
		}
	}
	return string(bs)
}

func main() {
	fmt.Println(ToLower("Hello! How are you?"))
}
