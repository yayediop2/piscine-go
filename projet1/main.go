package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func hexadecimal(s string) int {
	var n = 0
	for i := 0; i < len(s); i++ {
		n *= 16
		if s[i] >= '0' && s[i] <= '9' {
			n += int(s[i] - '0')
		} else if s[i] >= 'a' && s[i] <= 'z' {
			n += int(s[i] - 'a' + 10)
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			n += int(s[i] - 'A' + 10)
		}
	}
	return n
}

func binaire(s string) int {
	n := 0
	for _, l := range s {
		n *= 2
		if l == '0' || l == '1' {
			n += int(l - '0')
		}
	}
	return n
}

func ToUpper(s string) string {
	bs := []rune(s)
	for i, value := range bs {
		if value >= 97 && value <= 122 {
			bs[i] -= 32
		}
	}
	return string(bs)
}

/*
a, _ := strconv.ParseInt("1E", 16, 32)
//bin
b, _ := strconv.ParseInt("11", 2, 32)
*/

func ToLower(s string) string {
	bs := []rune(s)
	for i, value := range bs {
		if value >= 65 && value <= 90 {
			bs[i] += 32
		}
	}
	return string(bs)
}

func Capitalize(s string) string {
	sr := []rune(s)
	if len(s) > 0 {
		for i := 0; i < len(s); i++ {
			if sr[i] == ' ' {
				sr[i+1] = (sr[i+1] + ' ')
			}
		}
	}
	return string(sr)
}

func SplitWhiteSpaces(s string) []string {
	tab := []string{}
	m := ""
	for _, letter := range s {
		if letter != ' ' {
			m += string(letter)
		}
		if letter == ' ' || letter == rune(s[len(s)-1]) {
			if m != "" {
				tab = append(tab, m)
				m = "," // ENLEVER çA APRES -> ""
			}
		}
	}
	return tab
}

func recupF() {
	arg := os.Args[1:]
	for i := 0; i < len(arg); i++ {
		fichierla, _ := ioutil.ReadFile(arg[0])
		//fichierlabas, _ := ioutil.ReadFile(arg[1])
		str := string(fichierla)
		c := SplitWhiteSpaces(string(str))
		/* for i,l:=range c{
			fichierlabas[i]=byte(l)
		}
		fichierlabas=c */
		fmt.Print(c)

	}
}

func main() {
	recupF()
}
