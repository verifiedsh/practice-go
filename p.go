package main

import (
	"strconv"
	"strings"
)

func hexToDec(sentence string) string {
	words := strings.Fields(sentence)

	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" {
			word, err := strconv.ParseInt(words[i-1], 16, 64)
			if err != nil {
				continue
			}
			words[i-1] = strconv.FormatInt(word, 10)
			words = append(words[:i], words[i+1:]...)
		}
	}
	return strings.Join(words, " ")
}
