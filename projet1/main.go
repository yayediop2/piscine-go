package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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
				m = "" // ENLEVER çA APRES -> ""
			}
		}
	}
	return tab
}

func recupF() []string {
	arg := os.Args[1:]
	var o []string
	for i := 0; i < len(arg); i++ {
		fich1, _ := ioutil.ReadFile(arg[0])
		fich2 := ioutil.WriteFile(arg[1], []byte(strings.Join(o, " ")), 0777)
		if fich2 != nil {
			panic(fich2)
		}
		str := string(fich1)
		c := SplitWhiteSpaces(string(str))
		o = motcle(c)
		fmt.Print(o)
		//fmt.Print(c)
	}
	return o
}

func motcle(s []string) []string {
	//tab := []string{}
	for i := 0; i < len(s); i++ {
		var num int
		//hex
		if s[i] == "(hex)" {
			tab = append(tab, hexadecimal(s[i-1]))
		}
		//bin
		if s[i] == "(bin)" {
			tab = append(tab, binaire(s[i-1]))
		}
		//cap
		if s[i] == "(cap)" {
			tab = append(tab, Capitalize(s[i-1]))
		}
		// (cap,
		if s[i] == "(cap," {
			for _, l := range s[i+1] {
				if l >= '0' && l <= '9' {
					num = int(l - '0')
					break
				}
			}
			for p := len(s[i-1]); p >= num; p-- {
				tab = append(tab, Capitalize(s[i-num]))
				i++
			}
		}
		//low
		if s[i] == "(low)" {
			tab = append(tab, ToLower(s[i-1]))
		}
		// (low,
		if s[i] == "(low," {
			for _, l := range s[i+1] {
				if l >= '0' && l <= '9' {
					num = int(l - '0')
					break
				}
			}
			for p := len(s[i-1]); p >= num; p-- {
				tab = append(tab, ToLower(s[i-p]))
				i++
			}
		}
		//up
		if s[i] == "(up)" {
			tab = append(tab, ToUpper(s[i-1]))
		}
		// (up,
		if s[i] == "(up," {
			for _, l := range s[i+1] {
				if l >= '0' && l <= '9' {
					num = int(l - '0')
					break
				}
			}
			for p := len(s[i-1]); p >= num; p-- {
				tab = append(tab, ToUpper(s[i-num]))
				i++
			}
		}

	}
	return tab
}

func main() {
	recupF()
}

/*
func motcle(s []string) []string {
	//tab := []string{}
	for i := 0; i < len(s); i++ {
		var num int
		//hex
		if s[i] == "(hex)" {
			tab = append(tab, hexadecimal(s[i-1]))
		}
		//bin
		if s[i] == "(bin)" {
			tab = append(tab, binaire(s[i-1]))
		}
		//cap
		if s[i] == "(cap)" {
			tab = append(tab, Capitalize(s[i-1]))
		}
		// (cap,
		if s[i] == "(cap," {
			for _, l := range s[i+1] {
				if l >= '0' && l <= '9' {
					num = int(l - '0')
					break
				}
			}
			for p := len(s[i-1]); p >= num; p-- {
				tab = append(tab, Capitalize(s[i-num]))
				i++
			}
		}
		//low
		if s[i] == "(low)" {
			tab = append(tab, ToLower(s[i-1]))
		}
		// (low,
		if s[i] == "(low," {
			for _, l := range s[i+1] {
				if l >= '0' && l <= '9' {
					num = int(l - '0')
					break
				}
			}
			for p := len(s[i-1]); p >= num; p-- {
				tab = append(tab, ToLower(s[i-p]))
				i++
			}
		}
		//up
		if s[i] == "(up)" {
			tab = append(tab, ToUpper(s[i-1]))
		}
		// (up,
		if s[i] == "(up," {
			for _, l := range s[i+1] {
				if l >= '0' && l <= '9' {
					num = int(l - '0')
					break
				}
			}
			for p := len(s[i-1]); p >= num; p-- {
				tab = append(tab, ToUpper(s[i-num]))
				i++
			}
		}

	}
	return tab
}
*/
