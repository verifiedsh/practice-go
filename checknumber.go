package main

func CheckNumber(arg string) bool {
	for _, vrfd := range arg {
		if vrfd >= '0' && vrfd <= '9' {
			return true
		}
	}
	return false
}

// checknumber
// Instructions
// Write a function that takes a string as an argument and returns true if the string contains any number, otherwise return false.
// Expected function
// func CheckNumber(arg string) bool {
// }
// Usage
// Here is a possible program to test your function:
