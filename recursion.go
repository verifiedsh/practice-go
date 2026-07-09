package main

import (
	"fmt"
)

func testcount(x int) int {
	if x > 11 {
		return 0
	}
	fmt.Println(x)
	return testcount(x + 1)
}
