package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func LoadBanner(name string) (map[rune][]string, error) {
	result := map[rune][]string{}
	if !strings.Contains(name, ".txt") {
		name += ".txt"
	}
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("empty banner file")
	}
	fileLines := strings.Split(string(data), "\n")
	if len(fileLines) != 856 {
		return nil, errors.New("invalid file content")
	}
	for r := ' '; r <= '~'; r++ {
		index := int(r-' ') * 9
		result[r] = fileLines[index+1 : index+9]
	}
	return result, nil
}

func main() {
	fmt.Println(LoadBanner("standard.txt"))
	fmt.Println(LoadBanner("shadow"))
	fmt.Println(LoadBanner("thinkertoy"))
}
