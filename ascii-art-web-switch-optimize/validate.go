package main

import "fmt"

func Validateinput(text string) (rune, error) {
	for _, r := range text {
		if r != '\r' && r != '\n' && (r < ' ' || r > '~') {
			return r, fmt.Errorf("unsupported character: %v", r)
		}
	}
	return 0, nil
}
