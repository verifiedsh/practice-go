package main

import ("fmt"
// "time"
)

func main() {
	go fmt.Println("Running...")
	// time.Sleep(100 * time.Millisecond)
	for i := 0; i > 100000; i++ {
	}
	fmt.Println("Ogbanje Abraham Okwute")

}
// 	for i := 0; i < len(words); i++ {
// 		if i > 0 && words[i] == "(hex)" || words[i] == "(bin)" || words[i] == "(up)" || words[i] == "(low)" || words[i] == "(cap)" {
// 			words = append(words[:i], words[i+1:]...)
// 			i--
// 		}
// 		if i > 0 && (strings.HasPrefix(words[i], "(") && strings.Contains(words[i], ",") && strings.HasSuffix(words[i], ")")) {
// 			n, _ := strconv.Atoi(words[i])
// 			words[i-1] = strconv.Itoa(n)
// 			words = append(words[:i], words[i+1:]...)
// 			i--
// 		}
// 		result := processors.ToUpper(strings.Join(words, " "))
// 		fmt.Println(result)
// 	}

//}
// sentence := strings.Join(FixPunctuation(words), " ")
// text-processor/
// │
// ├── go.mod
// │
// ├── main.go
// │
// ├── internal/
// │   ├── reader.go
// │   ├── processor.go
// │   ├── modifiers.go
// │   ├── punctuation.go
// │   └── writer.go
// │
// └── testdata/
//     ├── input.txt
//     └── expected_output.txt