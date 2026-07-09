package main

import (
	"regexp"
	"strings"
)

func singleQ(s string) string {
	x := strings.Split(s, `'`)

	y := ""
	for i := 1; i < len(x); i++ {
		if i%2 == 1 {
			x[i] = strings.TrimSpace(x[i])
		}
	}

	y = strings.Join(x, `'`)

	z := strings.Split(y, `"`)

	for i := 1; i < len(z); i++ {
		if i%2 == 1 {
			z[i] = strings.TrimSpace(z[i])
		}
	}
	r := strings.Join(z, `"`)
	return strings.Join(strings.Fields(r), " ")

}

func Punc(s string) string {
	re := regexp.MustCompile(`\s+([,./;:!])`)
	s = re.ReplaceAllString(s, "$1")

	re2 := regexp.MustCompile(`([,./;:!])([^\s,./;:!])`)
	s = re2.ReplaceAllString(s, "$1 $2")

	return s
}
