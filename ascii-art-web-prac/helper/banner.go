package helper

import (
	"errors"
	"os"
)

func LoadBanner(banner string) ([]string, error) {
	data, err := os.ReadFile("banner/" + banner)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("empty banner file")
	}
	fileLines := SplitInput(string(data))
	if len(fileLines) != 856 {
		return nil, errors.New("invalid file content")
	}
	return fileLines, nil
}
