package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func hexadecimal(s string) string {
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
	str := strconv.Itoa(n)
	return string(str)
}

func binaire(s string) string {
	n := 0
	for _, l := range s {
		n *= 2
		if l == '0' || l == '1' {
			n += int(l - '0')
		}
	}
	str := strconv.Itoa(n)
	return string(str)
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
			if i == 0 && s[i] >= 'a' && s[i] <= 'z' {
				sr[i] = ((sr[i]) - ' ')
				/* if sr[i] == ' ' {
					sr[i+1] = (sr[i+1] + ' ')
				} */
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
		o := motcle(c)
		fmt.Print(o)
		fmt.Print(c)
	}
}

func motcle(s []string) []string {
	tab := []string{}
	for i := 0; i < len(s); i++ {
		var num int
		//cap
		if s[i] == "(cap)" {
			tab = append(tab, Capitalize(s[i-1]))
		}
		if s[i] == "(cap" {
			nomb := s[i+1]
			for _, l := range nomb {
				if l >= '0' && l <= '9' {
					num = int(l)
				}
			}
		}
		//hex
		if s[i] == "(hex)" {
			tab = append(tab, hexadecimal(s[i-1]))
		}
		//bin
		if s[i] == "(bin)" {
			tab = append(tab, binaire(s[i-1]))
		}
		//low
		if s[i] == "(low)" {
			tab = append(tab, ToLower(s[i-1]))
		}
		//up
		if s[i] == "(up)" {
			tab = append(tab, ToUpper(s[i-1]))
		}

	}
	return tab
}

func main() {
	recupF()
}
