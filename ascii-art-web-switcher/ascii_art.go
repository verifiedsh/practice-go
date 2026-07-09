package main

import(
	"errors"
	"os"
	"strings"
)

func LoadBanner(banner string) ([]string, error) {
	data, err := os.ReadFile(banner+".txt")
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("empty banner file")
	}
	lines := strings.Split(string(data), "\n")
	if len(lines) != 856 {
		return nil, errors.New("invalid file content")
	}
	return lines, nil
}

func GenerateArt(words []string, banner []string) string {
	var result strings.Builder
	for index, word := range words {
		for i := 1; i < 9; i++ {
			for _, char := range word {
				result.WriteString(banner[i+int(char-' ')*9])
			}
			if i < 8 {
				result.WriteString("\n")
			}
		}
		if index < len(words) {
			result.WriteString("\n")
		}
	}
	return result.String()
}