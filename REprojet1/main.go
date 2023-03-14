package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// J'ai remplacé tous mes 0 par +
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
			/* for _, l := range s[i+1] {
				if l >= '0' && l <= '9' {
					num = int(l - '0')
					break
				}
			} */
			nextS := s[i+1]
			num, _ = strconv.Atoi(nextS[:len(nextS)-1])
			if num <= len(s)-2 {
				s[i], s[i+1] = "", ""
				for p := num; p >= 0; p-- {
					s[i] = Capitalize(s[i])
					i--
				}
			} else {
				num = i
				s[i], s[i+1] = "", ""
				for p := num; p >= 0; p-- {
					s[i] = Capitalize(s[i])
					i--
				}
				break
			}
		}
		//low
		if s[i] == "(low)" {
			s[i-1] = ToLower(s[i-1])
			s[i] = ""
		}
		// (low,
		if s[i] == "(low," {
			nextS := s[i+1]
			num, _ = strconv.Atoi(nextS[:len(nextS)-1])
			if num <= len(s)-2 {
				s[i], s[i+1] = "", ""
				for p := num; p >= 0; p-- {
					s[i-num] = ToLower(s[i-num])
					i++
				}
			} else {
				num = i
				s[i], s[i+1] = "", ""
				for p := num; p >= 0; p-- {
					s[i] = ToLower(s[i])
					i--
				}
				break
			}

		}
		//up
		if s[i] == "(up)" {
			s[i-1] = ToUpper(s[i-1])
			s[i] = ""
		}
		// (up,
		if s[i] == "(up," {
			nextS := s[i+1]
			num, _ = strconv.Atoi(nextS[:len(nextS)-1])
			if num <= len(s)-2 {
				s[i], s[i+1] = "", ""
				for p := num; p >= 0; p-- {
					s[i-num] = ToUpper(s[i-num])
					i++
				}
			} else {
				num = i
				s[i], s[i+1] = "", ""
				for p := num; p >= 0; p-- {
					s[i] = ToUpper(s[i])
					i--
				}
				break
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
		if s[i] == "A" {
			if isVowel(string(s[i+1][0])) {
				s[i] = "An"
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

func Ponctuation(saa []string) []string {
	s := FromSliceStr2Rune(saa)
	for i := 0; i < len(s); i++ {
		if IsPonctuation(s[i]) {
			if s[i-1] == ' ' {
				if s[i] == '.' && i+2 > len(s)-1 && s[i-1] == ' ' { //i+2 > len(s)-1
					s[i-1] = '+' // pourquoi le i+1 qu'ils ont mis? ça ne marche meme paas
				} else { // manam If the current rune is not a period (.) at the end of the string, it swaps
					s[i-1] = s[i] // ça swap ici hein
					s[i] = ' '
				}
			} else if i+1 < len(s)-2 && s[i+1] != ' ' && s[i-1] != '.' {
				s[i] += s[i] + ' '
			}
		}
	}
	sss := []rune{}
	for _, l := range s {
		if l != '+' {
			sss = append(sss, l)
		}
	}
	sx := string(sss)
	slice := strings.Split(sx, " ")
	return slice
}

/* func IsLetter(s string) bool {
	bs := []byte(s)
	for _, value := range bs {
		if value >= 97 && value <= 122 || value >= 65 && value <= 90 {
			return false
		}
	}
	return true
} */

func FromSliceStr2Rune(saa []string) []rune {
	v := espacesSup(saa)
	sa := strings.Join(v, " ")
	s := []rune(sa)
	return s
}

func espacesSup(s []string) []string {
	var tableau []string
	for _, l := range s {
		if l != "" && l != "0" {
			tableau = append(tableau, l)
		}
	}
	return tableau
}

func IsPonctuation(s rune) bool {
	if s == '.' || s == ',' || s == '!' || s == '?' || s == ':' || s == ';' {
		return true
	}
	return false
}

func apostorophiii(sa []string) []string {
	s := FromSliceStr2Rune(sa)
	A := 0
	if len(s) > 1 { // pour gérer le cas ou on a une seule apostrophe
		for i := 0; i < len(s); i++ {
			if s[i] == '\'' && A == 0 {
				if s[i+1] == ' ' { // je ne pense paaas qu'on ait besoin du len là
					s[i+1] = '+'
				}
				A++
				continue
			}
			if s[i] == '\'' && A == 1 {
				if s[i-1] == ' ' {
					s[i-1] = '+'
				}
				A = 0
			}
			if s[len(s)-1] == '\'' {
				if s[len(s)-2] == ' ' {
					s[len(s)-2] = '+'
				}
			}
		}
	}

	var tableau []rune
	for _, l := range s {
		if l != '+' {
			tableau = append(tableau, l)
		}
	}
	sx := string(tableau)
	slice := strings.Split(sx, " ")
	return slice
}

func ecrireF(cv []string) {
	arg := os.Args[1:]
	fich2 := ioutil.WriteFile(arg[1], []byte(strings.Join(cv, " ")), 0777) // join saute
	if fich2 != nil {
		panic(fich2)
	}
}

func main() {
	a := recupF()
	b := motcle(a)
	c := AToAn(b)
	d := Ponctuation(c)
	e := espacesSup(d)
	tableau := apostorophiii(e)
	ecrireF(tableau)
}

/*
func apostorophiii(sa []string) []string {
	s := FromSliceStr2Rune(sa)
	A := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\'' && A == 0 {
			if s[i+1] == ' ' {
				s[i+1] = s[i] // ça swap ici hein
				s[i] = ' '
			}
			A++
			continue
		}
		if s[i] == '\'' && A == 1 {
			if s[i-1] == ' ' {
				s[i-1] = s[i]
				s[i] = ' '
			}
			A = 0
		}
	}
	sx := string(s)
	slice := strings.Split(sx, " ")
	return slice
}*/
