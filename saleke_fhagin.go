package main

import (
	"strings"
	"unicode"
)

func Transform(s string) []string {
	new := strings.FieldsFunc(s, func(r rune) bool {
		return unicode.IsPunct(r) || r == ' ' || r == 'd' || r == 'v'
	})
	return new
	// r := []rune(s)

	// for i := 0; i < len(r); i++ {
	// 	char := r[i]
	// 	if (char >= 33 && char <= 47) ||
	// 		(char >= 58 && char <= 64) ||
	// 		(char >= 91 && char <= 96) ||
	// 		(char >= 123 && char <= 126) {
	// 		r[i] = ' '
	// 	}
	// }
	// return strings.Fields(string(r))
}
