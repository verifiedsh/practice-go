package main

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(bannerFile string) ([]string, error) {
	data, err := os.ReadFile(bannerFile)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("empty banner file")
	}
	lines := strings.Split(string(data), "\n")
	if len(lines) != 856 {
		return nil, errors.New("invalid file content")
	}
	return lines, nil
}
