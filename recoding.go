package main

import (
	"strings"
)

func UpperCaseLastN(s []string, n int) []string {
	for i := len(s) - n; i < len(s); i++ {
		s[i] = strings.ToUpper(s[i])
	}
	return s
}
