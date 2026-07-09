package main

import "fmt"

func main() {
	x := []int{15, -4, 56, 10, -23}
	y := []int{14, -9, 56, 14, -23}
	fmt.Println(recursion(x, y))
}

func absDiff(a, b []int) int {

	diff := 0
	for i := range len(a) {
		dif := (a[i] - b[i])
		if dif < 0 {
			diff += -dif
		} else {
			diff += dif
		}
	}
	return diff
}

func recursion(c, d []int) int {
	led := c[len(c)-1] - d[len(d)-1]
	if led < 0 {
		led = -led
	}
	diff := absDiff(c[:len(c)-1], d[:len(d)-1]) + led
	return diff
}
