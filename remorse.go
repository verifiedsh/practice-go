package main

import (
	"fmt"
	"strings"
)

func hello() string {
	fmt.Println("What is your first name?")
	var name string
	fmt.Scanf("%v", &name)
	for i := 0; i < len(name); i++ {
		if string(name[0]) >= "a" && string(name[0]) <= "z" || string(name[0]) >= "A" && string(name[0]) <= "Z" {
			name = strings.ToUpper(string(name[0])) + name[1:]
			fmt.Println(`hello ` + name)
		} else if string(name[i]) >= "0" && string(name[i]) <= "9" || strings.ContainsAny(string(name[i]), ",.;;!?") {
			fmt.Println("Enter a valid name with letters not numbers or punctuation")
			hello()
		} else if !(string(name[0]) >= "a" && string(name[0]) <= "z") || !(string(name[0]) >= "A" && string(name[0]) <= "Z") || !strings.ContainsAny(string(name[i]), ",.;;!?") {
			fmt.Println("Enter a valid name with letters")
			hello()
		}
	}
	return name
}
