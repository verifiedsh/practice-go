package helper

import (
	"fmt"
)

func ValidateInput(text string) ([]rune, error) {
	var odds []rune
	for _, v := range text {
		if v != '\n' && v != '\r' && (v < ' ' || v > '~') {
			odds = append(odds, v)
			continue
		}
	}
	if odds != nil {
		return odds, fmt.Errorf("unsupported character(s): %c", odds)
	}
	return nil, nil
}
