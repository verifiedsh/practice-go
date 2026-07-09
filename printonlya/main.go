// printonlya

package main

// import "github.com/01-edu/z01"

// func main() {
// 	z01.PrintRune('a')
// 	zo1.PrintRune('\n')
// }

// package main

import "fmt"

// func CheckNumber(arg string) bool {
// 	for _, i := range arg {
// 		if i >= '0' && i <= '9' {
// 			return true
// 		}
// 	}
// 	return false
// }

// func main() {
// 	fmt.Println(CheckNumber("Hello"))
// 	fmt.Println(CheckNumber("Hello1"))
// }

// PrintIf

// package main

// import "fmt"

// func PrintIf(str string) string {
// 	if str == "" {
// 		return "G\n"
// 	}
// 	if len(str) >= 3 {
// 		return "G\n"
// 	}
// 	return "Invalid Input\n"
// }

// func main() {
// 	fmt.Print(PrintIf("abcdefz"))
// 	fmt.Print(PrintIf("abc"))
// 	fmt.Print(PrintIf(""))
// 	fmt.Print(PrintIf("14"))
// }

// package main

// import "fmt"

// func CountAlpha(s string) int {
// 	count := 0
// 	for _, i := range s {
// 		if (i >= 'A' && i <= 'Z') || (i >= 'a' && i <= 'z') {
// 			count++
// 		}
// 	}
// 	return count
// }

// func main() {
// 	fmt.Println(CountAlpha("Hello world"))
// 	fmt.Println(CountAlpha("H e l l o"))
// 	fmt.Println(CountAlpha("H1e2l3l4o"))
// }

// var array_name = [length]datatype{values} //  Here length is specified
// var array_name =[...]datatype{values} //  Here length is inferred

// array_name := [length]datatype{values} //  Here length is specified
// array_name :=[...]datatype{values} //  Here length is inferred

// array and slice practice

// package main

// import "fmt"

// func main() {
// var arr1 = [5]int{5, 4, 9, 2, 1}
// myslice1 := arr1[2:3]

// fmt.Printf("myslice = %v\n", (myslice1))
// fmt.Printf("length = %d\n", len(myslice1))
// fmt.Printf("capacity = %d\n", cap(myslice1))

// 	myslice1 := make([]int, 5, 10)
// 	fmt.Printf("myslice = %v\n", myslice1)
// 	fmt.Printf("length = %d\n", len(myslice1))
// 	fmt.Printf("capacity = %d\n", cap(myslice1))
// }

// fmt.Println(arr1)
// 	fmt.Println(arr2)
// arr2 := [3]string{"Ben", "Paul", "John"}

//slice practice

// package main

// import "fmt"

func main() {
	mySlice1 := []int{1, 2, 3, 4}
	fmt.Println(len(mySlice1))
	fmt.Println(cap(mySlice1))
	fmt.Println(mySlice1)
}
