package main

import (
	"fmt"
	"strings"
)

func StringToArt(input string) string {
	numbers := map[rune][]string{
		'0': {
			"  _  ",
			" | | ",
			" | | ",
			" | | ",
			" |_| ",
		},
		'1': {
			"     ",
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
		},
		'2': {
			"  _  ",
			" | | ",
			"   | ",
			"  /  ",
			" |__ ",
		},
		'3': {
			"  _  ",
			" | | ",
			"  _| ",
			"   | ",
			" |_| ",
		},
		'4': {
			"     ",
			" | | ",
			" |_|_",
			"   | ",
			"   | ",
		},
		'5': {
			"  _  ",
			" |   ",
			" |_  ",
			"   | ",
			"  _| ",
		},
		'6': {
			"  _  ",
			" |   ",
			" |_  ",
			" | | ",
			" |_| ",
		},
		'7': {
			"  _  ",
			"   | ",
			"   | ",
			"   | ",
			"   | ",
		},
		'8': {
			"  _  ",
			" | | ",
			" |_| ",
			" | | ",
			" |_| ",
		},
		'9': {
			"  _  ",
			" | | ",
			" |_| ",
			"   | ",
			"   | ",
		},
	}
	joiner := []string{}
	var result strings.Builder
	nums := strings.Split(input, "\n")
	fmt.Println(nums)
	for _, num := range nums {
		for i := range 5 {
			for _, n := range num {
				result.WriteString(numbers[n][i])
			}
			joiner = append(joiner, result.String())
			result.Reset()
		}
	}
	return strings.Join(joiner, "\n")
}

func main() {
	fmt.Println(StringToArt("12323\n45645"))
}

/*
1. Create a function StringToArt that converts a string into ASCII art using only pipe |, underscore _, forward slash /, backslash \, and space characters. The output should look like simple line art.

Requirements:

Function signature: func StringToArt(input string) string

Support digits 0-9 only

Each digit should be 5 characters wide and 5 lines tall

Digits should be joined without spacing

Handle multiple lines in input separated by \n

Return empty string for invalid input

File Structure:

ascii-art/
├── main.go
├── converter/
│   ├── converter.go
│   └── converter_test.go

*/
