package main

import "fmt"

func ToUpper(s string) string {
	bs := []byte(s)
	for _, value := range bs {
		if value >= 97 && value <= 122 {
			value - 32
		}
	}
	return string(bs)
}

func main() {
	fmt.Println(ToUpper("Hello! How are you?"))
}
