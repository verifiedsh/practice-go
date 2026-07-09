package main

import (
	"fmt"
)

func prob() {
	s := []string{"c", "d", "e", "n", "n", "c", "k", "m", "r", "a"}

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if i != len(s)-1 && s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
			if s[i] == s[j] {
				s = append(s[:i], s[i+1:]...)
			}
		}
	}
	fmt.Println(s)
}
