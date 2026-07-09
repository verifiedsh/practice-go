/*
Write a function that:
1. Write a function that returns the last letter of a word.
2. Write a function that returns the first two letters of a word.
3. Write a function that takes a slice of words and applies uppercase to the first word.
4. Write a function that converts 10 string to 10 integer.
5. Write a function that joins 2 strings without space.
6. Write a function that checks if 2 strings are the same.
*/
// package main

// import(
// "fmt"
// "strconv"
// "strings"
// )
// func lastL(s string) string {
// 	return strings.ToUpper(string(s[len(s)-1]))
// }
// func firstTwoL(s string) string {
// 	return s[0:2]
// }

// func upperCaseFirstW(words []string) []string {
// 	words[0] = strings.ToUpper(words[0])
// 	return words
// }

// func convToInt(s string) int {
// 	r, _ := strconv.Atoi(s)
// 	return r
// }

// func joinTwoW(a, b string) string {
// 	return strings.ToUpper(a)+b
// }

// func checkStr(a, b string) string {
// 	if a == b {
// 		return "TRUE"
// 	} else {
// 		return "FALSE"
// 	}
// }

// func main() {
// fmt.Println(lastL("mentor"))
// fmt.Println(firstTwoL("mentor"))
// fmt.Println(upperCaseFirstW([]string{"banana", "mango", "orange"}))
// fmt.Println(convToInt("10"))
// fmt.Println(joinTwoW("leef", "center"))
// fmt.Println(checkStr("hi", "hi"))
// fmt.Println(checkStr("hi", "hello"))
// }

// write a program that prints all the numbers from 1 to 100 that are evenly divisible by 3.
package main

import "fmt"

func off() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
