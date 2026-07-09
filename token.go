package main

import (
	"unicode"
)

func Tokenize(text string) []string {

	tokens := []string{}

	currenttoken := ""

	for _, s := range text {

		if unicode.IsLetter(s) || unicode.IsDigit(s) {

			currenttoken += string(s)

		} else if unicode.IsSpace(s) {

			if currenttoken != "" {

				tokens = append(tokens, currenttoken)

				currenttoken = ""
			}

		} else {
			if currenttoken != "" {
				tokens = append(tokens, currenttoken)
				currenttoken = ""
			}
			tokens = append(tokens, string(s))
		}

	}
	if currenttoken != "" {

		tokens = append(tokens, currenttoken)

	}
	return tokens
	// for i := 0; i < len(tokens); i++ {
	// 	if strings.ContainsAny(tokens[i], ",.;:!?") {
	// 		tokens[i-1] = tokens[i-1] + tokens[i]
	// 		fmt.Println(tokens)

	// 	}
	// 	tokens = append(tokens[:i], tokens[i+1:]...)
	// }
	// return tokens
}
