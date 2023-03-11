package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func ToUpper(s string) string {
	bs := []rune(s)
	for i, value := range bs {
		if value >= 97 && value <= 122 {
			bs[i] -= 32
		}
	}
	return string(bs)
}

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
			}
		}
	}
	return string(sr)
}

func recupF() []string {
	arg := os.Args[1:]
	var c []string
	for i := 0; i < len(arg); i++ {
		fich1, _ := ioutil.ReadFile(arg[0])

		str := string(fich1)
		c = strings.Fields(string(str)) // splitwhitespaces saute
	}
	return c
}

func motcle(s []string) []string {
	for i := 0; i < len(s); i++ {
		var num int

		//hex
		if s[i] == "(hex)" {
			a, _ := strconv.ParseInt(s[i-1], 16, 32)
			b := strconv.Itoa(int(a))
			s[i-1] = b
			s[i] = ""
		}
		//bin
		if s[i] == "(bin)" {
			a, _ := strconv.ParseInt(s[i-1], 2, 32)
			b := strconv.Itoa(int(a))
			s[i-1] = b
			s[i] = ""
		}
		//cap
		if s[i] == "(cap)" {
			s[i-1] = Capitalize(s[i-1])
			s[i] = ""
		}
		// (cap,
		if s[i] == "(cap," {
			for _, l := range s[i+1] {
				if l >= '0' && l <= '9' {
					num = int(l - '0')
					break
				}
			}
			s[i], s[i+1] = "", ""
			for p := 1; p <= num; p++ {
				s[i-num] = Capitalize(s[i-num])
				i++
			}
		}
		//s = append(s[:i], s[i+2:]...) // khady dit de faire tous les traitements puis de le mettre hors de la boucle. j'ai besoin d'avoir tous les index
		//low
		if s[i] == "(low)" {
			s[i-1] = ToLower(s[i-1])
			s[i] = ""
		}
		// (low,
		if s[i] == "(low," {
			for _, l := range s[i+1] {
				if l >= '0' && l <= '9' { // gerer si on 14 par exemple. Eviter le out of range
					num = int(l - '0')
					break // on aurait plus besoin de ça
				}
			}
			s[i], s[i+1] = "", ""
			for p := 1; p <= num; p++ {
				s[i-num] = ToLower(s[i-num])
				i++
			}
		}
		//up
		if s[i] == "(up)" {
			s[i-1] = ToUpper(s[i-1])
			s[i] = ""
		}
		// (up,
		if s[i] == "(up," {
			for _, l := range s[i+1] {
				if l >= '0' && l <= '9' {
					num = int(l - '0')
					break
				}
			}
			s[i], s[i+1] = "", ""
			for p := 1; p <= num; p++ {
				s[i-num] = ToUpper(s[i-num])
				i++
			}
		}
	}
	return s
}

func AToAn(s []string) []string {
	for i := 0; i < len(s); i++ {
		if s[i] == "a" {
			if isVowel(string(s[i+1][0])) {
				s[i] = "an"
			}
		}
	}
	return s
}

func isVowel(s string) bool {
	v := []string{"a", "e", "u", "i", "o", "A", "E", "U", "I", "O"}
	for _, l := range v {
		if s == l {
			return true
		}
	}
	return false
}

func ecrireF(cv []string) {
	arg := os.Args[1:]
	fich2 := ioutil.WriteFile(arg[1], []byte(strings.Join(cv, " ")), 0777) // join saute
	if fich2 != nil {
		panic(fich2)
	}
}

func main() {
	var tableau []string
	a := recupF()
	b := motcle(a)
	c := AToAn(b)
	for _, l := range c {
		if l != "" {
			tableau = append(tableau, l) // ICI LE TABLEAU N'EST LIÉ QU'A ATOAN
		}
	}
	ecrireF(tableau)
}
