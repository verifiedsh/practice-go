package main

import (
	"strings"
)

func punct(sentence string) string {
	words := strings.Fields(sentence)
	punc := ",.;:!?"

	newWords := []string{}

	for _, v := range words {
		if len(v) > 1 && strings.ContainsAny(string(v[0]), punc) {
			newWords = append(newWords, string(v[0]), v[1:])
		} else {
			newWords = append(newWords, v)
		}
	}

	sorted := []string{}

	for i := 0; i < len(newWords); i++ {
		if len(sorted) > 0 && strings.ContainsAny(newWords[i], punc) {
			sorted[len(sorted)-1] = sorted[len(sorted)-1] + newWords[i]
		} else {
			sorted = append(sorted, newWords[i])
		}
	}

	// if len(words[i]) == 1 && strings.ContainsAny(words[i], punc) {
	// 			words[i-1] += words[i]
	// 			words = append(words[:i], words[i+1:]...)
	// 		}

	return strings.Join(sorted, " ")
}

/*
	[]string{"Hello", ",", "world", ".", ".", "."}

	[]string{"Hello", ",world"}

	[]string{"Hello", ",", "world"}
*/
