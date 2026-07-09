package helper

import "strings"

func GenerateArt(words []string, banner []string) string {
	var result strings.Builder
	for index, word := range words {
		for i := 1; i < 9; i++ {
			for _, char := range word {
				result.WriteString(banner[i+int(char-' ')*9])
			}
			if i != 8 {
				result.WriteString("\n")
			}
		}
		if index != len(words)-1 {
			result.WriteString("\n")
		}
	}
	return result.String()
}
