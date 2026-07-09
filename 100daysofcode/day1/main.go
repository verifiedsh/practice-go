package main

import "fmt"

func countLowNum(num []int) []int {
	count := 0
	less := []int{}

	for _, i := range num {
		for _, j := range num {
			if i > j {
				count++
			}
		}
		less = append(less, count)
		count = 0
	}
	return less
}

func main() {
	nums := []int{8, 3, 2, 5, 6}

	fmt.Println(countLowNum(nums))
}
