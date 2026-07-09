package main

import (
	"fmt"
	"strings"
)

func tutorial() {
	word := "hehloiiii"
	count := 0
	var evenStr string
	if len(word)%2 == 0 {
		fmt.Println(word[len(word)/2-1 : len(word)/2+1])
	} else {
		fmt.Println(string(word[len(word)/2]))
	}
	for i := 0; i < len(word); i++ {
		for j := 0; j < len(word); j++ {
			if word[i] == word[j] {
				count++
			}
		}
		if count%2 == 0 && !strings.ContainsAny(evenStr, string(word[i])) {
			evenStr += string(word[i])
			count = 0
		} else if count%2 != 0 {
			count = 0
		}
	}
	fmt.Println(evenStr)
}
