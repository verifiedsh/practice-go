package main

import (
	"fmt"
	"strings"
	"unicode"
)

// 1. Reverse ASCII string
func Reverse(s string) string {
	var x string
	for i := len(s) - 1; i >= 0; i-- {
		x += string(s[i])
	}
	return x
}

// 2. Count lowercase vowels
func CountVowels(s string) int {
	var count int
	lowVowels := "aeiou"
	for _, v := range s {
		if strings.ContainsAny(string(v), lowVowels) {
			count++
		}
	}
	return count
}

// 3. Replace all spaces with underscores
func UnderscoreSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "_")
}

// 4. Capitalize first letter of each word
func CapitalizeWords(s string) string {
	x := strings.Fields(s)
	for i := 0; i < len(x); i++ {
		x[i] = strings.ToUpper(x[i][:1]) + strings.ToLower(x[i][1:])
	}
	return strings.Join(x, " ")
}

// 5. Count how many times a letter occurs
func CountLetter(s string, target rune) int {
	var count int
	for _, v := range s {
		if v == rune(target) {
			count++
		}
	}
	return count
}

// 🟡 LEVEL 2 — MEDIUM LOGIC & LOOPING
// 6. Remove all digits
func RemoveDigits(s string) string {
	s = strings.ReplaceAll(s, ",", " ")
	var x string
	xx := []string{}
	ss := strings.Fields(s)
	for i := 0; i < len(ss); i++ {
		for _, v := range ss[i] {
			if unicode.IsLetter(v) {
				x += string(v)
			}
		}
		xx = append(xx, x)
		x = ""
	}

	return strings.Join(xx, " ")
}

// 7. Extract last word (without Split)
func LastWord(s string) string {
	i := len(s) - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}
	end := i

	for i >= 0 && s[i] != ' ' {
		i--
	}
	start := i + 1
	return s[start : end+1]

	// s = strings.TrimRight(s, "\n\r\t")
	// for i := len(s) - 1; i >= 0; i-- {
	// 	if s[i] == ' ' || s[i] == '\n' || s[i] == '\t' || s[i] == 'r' {
	// 		return s[i+1:]
	// 	}
	// }
	// return s
}

// 8. Toggle case
func ToggleCase(s string) string {
	var x []rune
	for _, v := range s {
		if unicode.IsLower(v) {
			x = append(x, unicode.ToUpper(v))
		} else if unicode.IsUpper(v) {
			x = append(x, unicode.ToLower(v))
		}
	}
	return string(x)
}

// 9. Remove duplicate characters (first occurrence only)
func UniqueChars(s string) string {
	var x string
	for i := 0; i < len(s); i++ {
		if !strings.ContainsAny(x, string(s[i])) {
			x += string(s[i])
		} else {
			continue
		}
	}
	return x
}

// 10. Repeat each character N times
func ExpandChars(s string, n int) string {
	var x string
	for i := 0; i < len(s); i++ {
		for j := 0; j < n; j++ {
			x += string(s[i])
		}
	}
	return x
}

// 11. Reverse only the vowels
func ReverseVowels(s string) string {
	vowels := "aeiouAEIOU"
	x := ""
	j := len(s) - 1
	for i := 0; i < len(s); i++ {
		if !strings.ContainsAny(string(s[i]), vowels) {
			x += string(s[i])
		} else if strings.ContainsAny(string(s[i]), vowels) {
			for ; j >= 0; j-- {
				if strings.ContainsAny(string(s[j]), vowels) {
					x += string(s[j])
					j--
					break
				}
			}
		}
	}
	return x
}

// 12. Remove all extra spaces (leave only one between words)
func NormalizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// 13. Check if string is palindrome ignoring spaces
func CleanPalindrome(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	var x string
	for i := len(s) - 1; i >= 0; i-- {
		x += string(s[i])
	}
	fmt.Println(x)
	return x == s
}

// 14. Return the most frequent character
func MostFrequentChar(s string) string {
	y := make(map[string]int)
	count := 0
	x := ""
	for _, v := range s {
		y[string(v)]++
		if y[string(v)] > count {
			count = y[string(v)]
			x = string(v)
		} else if y[string(v)] >= count && !strings.ContainsAny(x, string(v)) {
			x += string(v)
		}
	}
	return x
}

// 15. Merge two strings character-by-character
func MergeString(a, b string) string {
	x := ""
	i := 0
	for v := 0; v < len(a); v++ {
		x += string(a[v])
		for ; i < len(b); i++ {
			x += string(b[i])
			i++
			break
		}
	}
	x += b[i:]
	return x
}

// 16. Reverse a Unicode string
func ReverseUnicode(s string) string {
	x := []rune(s)
	if len(s) == 1 {
		return s
	}
	return ReverseUnicode(string(x[1:])) + string(x[0])
}

// 18. Check if two strings are anagrams
// (ignore spaces, ignore case)
func Anagram(a, b string) bool {
	a, b = strings.ToLower(strings.ReplaceAll(a, "", " ")), strings.ToLower(strings.ReplaceAll(b, "", " "))

	aa, bb := strings.Fields(a), strings.Fields(b)
	for i := 0; i < len(aa); i++ {
		for j := i + 1; j < len(aa); j++ {
			if aa[i] < aa[j] {
				aa[i], aa[j] = aa[j], aa[i]
			}
			if bb[i] < bb[j] {
				bb[i], bb[j] = bb[j], bb[i]
			}
		}
	}
	a, b = strings.Join(aa, " "), strings.Join(bb, " ")
	return strings.Contains(a, b)
}
