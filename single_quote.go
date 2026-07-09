package main

import (
	"strings"
)

//	func fixSingleQuotes(word string) string {
//		var flag int
//		words := strings.Fields(word)
//		for index, value := range words {
//			if value == "'" && flag == 0 {
//				words[index+1] = value + words[index+1]
//				words[index] = ""
//				flag = 1
//			}
//			if value == "'" && flag == 1 {
//				words[index-1] = words[index-1] + value
//				words[index] = ""
//				flag = 0
//			}
//		}
//		result := strings.Join(words, " ")
//		result = strings.ReplaceAll(result, "  ", " ")
//		return result
//	}
func rev(s string) string {
	if s == "" {
		return s
	}
	return rev(s[1:]) + string(s[0])
}
func FixSingleQuotes(word string) []string {
	words := strings.Fields(word)
	var flag int
	for index, value := range words {
		if value == "'" && flag == 0 {
			words[index+1] = value + words[index+1]
			words[index] = ""
			flag = 1
		} else if value == "'" && flag == 1 {
			words[index-1] = words[index-1] + value
			words[index] = ""
			flag = 0
		}
	}
	result := strings.Join(words, " ")
	result = strings.ReplaceAll(result, "  ", " ")
	res := strings.Fields(result)
	return res
}
