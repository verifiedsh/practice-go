package main

import (
	"strconv"
	"strings"
)

func JoinModifiers(words []string) []string {
	result := []string{}
	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], "(up,") ||
			strings.HasPrefix(words[i], "(low,") ||
			strings.HasPrefix(words[i], "(cap") &&
				!strings.HasSuffix(words[i], ")") {
			if i+1 < len(words) && strings.HasSuffix(words[i+1], ")") {
				merged := strings.ReplaceAll(words[i]+words[i+1], " ", "")
				result = append(result, merged)
				i++
				continue
			}
		}
		result = append(result, words[i])
	}
	return result
}

func ProcessCase(words []string) []string {
	words = JoinModifiers(words)
	for i := 0; i < len(words); i++ {
		if i > 0 && words[i] == "(up)" {
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		} else if strings.HasPrefix(words[i], "(up,") && strings.HasSuffix(words[i], ")") {
			modifier := strings.TrimSpace(words[i][4 : len(words)-1])
			n, err := strconv.Atoi(modifier)
			if err == nil && n > 0 {
				for j := 1; j <= n && i-j >= 0; j++ {
					words[i-j] = strings.ToUpper(words[i-j])
				}
			}
			words = append(words[:i], words[i+1:]...)
			i--
		}

		if i > 0 && words[i] == "(low)" {
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		} else if strings.HasPrefix(words[i], "(low,") && strings.HasSuffix(words[i], ")") {
			modifiers := strings.TrimSpace(words[i][5 : len(words)-1])
			n, err := strconv.Atoi(modifiers)
			if err == nil && n > 0 {
				for j := 1; j <= n && i-j >= 0; j++ {
					words[i-j] = strings.ToLower(words[i-j])
				}
			}
			words = append(words[:i], words[i+1:]...)
			i--
		}

		if i > 0 && words[i] == "(cap)" {
			words[i-1] = strings.ToUpper(words[i-1][:1]) + strings.ToLower(words[i-1][1:])
			words = append(words[:i], words[i+1:]...)
			i--
		} else if strings.HasPrefix(words[i], "(cap,") && strings.HasSuffix(words[i], ")") {
			modifier := strings.TrimSpace(words[i][5 : len(words[i])-1])
			n, err := strconv.Atoi(modifier)
			if err == nil && n > 0 {
				for j := 1; j <= n && i-j >= 0; j++ {
					words[i-j] = strings.ToUpper(words[i-j][:1]) + strings.ToLower(words[i-j][1:])
				}
			}
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return words
}

func ProcessBase(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" && i > 0 {
			num, err := strconv.ParseInt(words[i-1], 16, 64)
			if err != nil {
				continue
			}
			words[i-1] = strconv.FormatInt(num, 10)
			words = append(words[:i], words[i+1:]...)
			i--
		}
		if words[i] == "(bin)" && i > 0 {
			num, err := strconv.ParseInt(words[i-1], 2, 64)
			if err != nil {
				continue
			}
			words[i-1] = strconv.FormatInt(num, 10)
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return words
}
