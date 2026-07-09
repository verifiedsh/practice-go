package main

import "strings"

func GenerateArt(rows []string, words []string) string {
	var result strings.Builder
	for index, word := range words {
		for i := 1; i < 9; i++ {
			for _, char := range word {
				result.WriteString(rows[i+int(char-' ')*9])
			}
			if i < 8 {
				result.WriteString("\n")
			}
		}
		if index < len(words)-1 {
			result.WriteString("\n")
		}
	}
	return result.String()
}
