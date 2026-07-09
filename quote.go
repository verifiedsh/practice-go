package main

import (
	"fmt"
	"strings"
	"unicode"
)

func quote(str string) string {
	// str = strings.Join(strings.Fields(str), " ")
	str = strings.TrimSpace(str)
	s := []string{}
	x := ""
	var indicator int

	for _, v := range str {
		if unicode.IsDigit(v) || unicode.IsLetter(v) {
			x += string(v)
		} else if len(x) > 0 && unicode.IsSpace(v) {
			s = append(s, x)
			x = ""
		} else if len(x) == 0 && strings.ContainsRune("'", v) && indicator == 0 {
			x += string(v)
			s = append(s, x)
			x = ""
			indicator = 1
		} else if strings.ContainsRune("'", v) && indicator == 1 {
			x += string(v)
			s = append(s, x)
			x = ""
			indicator = 0
		} else if unicode.IsPunct(v) {
			x += string(v)
			s = append(s, x)
			x = ""
		}
	}

	for i, v := range s {
		if v == "'" && indicator == 0 {
			s[i+1] = v + s[i+1]
			s[i] = ""
			indicator = 1
			s = append(s[:i], s[i+1:]...)
		} else if v == "'" && indicator == 1 {
			s[i-1] = s[i-1] + v
			s[i] = ""
			indicator = 0
			s = append(s[:i], s[i+1:]...)
		}
	}
	return strings.Join(s, " ")
}

func rmDuplicate(s string) string {
	x := strings.Fields(s)
	y := []string{}
	for i := 0; i < len(x); i++ {
		if strings.Count(strings.Join(y, " "), x[i]) < 1 {
			y = append(y, x[i])
		}
	}
	fmt.Printf("%q\n", strings.Join(y, " "))
	return strings.Join(y, " ")
}

// 15. Merge two strings character-by-character
func MergeStrings(a, b string) string {
	merged := ""
	x := 0
	for v := 0; v < len(a); v++ {
		merged += string(a[v])
		for ; x < len(b); x++ {
			merged += string(b[x])
			x++
			break

		}
	}
	merged += b[x:]
	return merged
}
