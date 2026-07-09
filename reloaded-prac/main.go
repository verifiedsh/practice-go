package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hexToDecimal(hexStr string) (int64, error) {
	return strconv.ParseInt(hexStr, 16, 64)
}

func binToDecimal(binStr string) (int64, error) {
	return strconv.ParseInt(binStr, 2, 64)
}

func capitalize(word string) string {
	return strings.ToUpper(string(word[0])) + strings.ToLower(string(word[1:]))
}

func UpperCaseLastN(words []string, n int) []string {
	for i := len(words) - n; i < len(words); i++ {
		words[i] = strings.ToUpper(words[i])
	}
	return words
}

func isPunc(word string) bool {
	for _, value := range word {
		if strings.ContainsRune("!,.?/;:'", value) {
			return true
		}
	}
	return false
}

func fixPunctuation(word string) string {
	word = strings.Join(strings.Fields(word), "")
	return strings.ReplaceAll(word, ",", ", ")

}

func fixSingleQuotes(word string) string {
	var flag int
	words := strings.Fields(word)
	for index, value := range words {
		if value == "'" && flag == 0 {
			words[index+1] = value + words[index+1]
			words[index] = ""
			flag = 1
		} else if value == "'" && flag == 1 {
			words[index-1] = words[index-1] + value
			words[index] = ""
			flag = 0
		}
	}
	result := strings.Join(words, " ")
	result = strings.ReplaceAll(result, "  ", " ")
	return result
}

func aOrAn(word string) string {
	if strings.ContainsAny(string(word[0]), "aeiouhAEIOUH") {
		return "An " + word
	}
	return "A " + word
}

func fixArticles(words string) string {
	word := strings.Fields(words)
	for i := 0; i < len(word)-1; i++ {
		next := word[i+1]
		if len(word) == 0 {
			continue
		}
		if word[i] == "a" && strings.ContainsAny(string(next[0]), "aeiouhAEIOUH") {
			word[i] = "an"
		} else if word[i] == "A" && strings.ContainsAny(string(next[0]), "aeiouhAEIOUH") {
			word[i] = "An"
		}
	}
	result := strings.Join(word, " ")
	return result
}

// func FixArticles(words []string) []string {
// 	vowels := "aeiouhAEIOUH"
// 	for i := 0; i < len(words)-1; i++ {
// 		next := words[i+1]
// 		if len(words) == 0 {
// 			continue
// 		}
// 		if words[i] == "a" && strings.ContainsRune(vowels, rune(next[0])) {
// 			words[i] = "an"
// 		} else if words[i] == "A" && strings.ContainsRune(vowels, rune(next[0])) {
// 			words[i] = "An"
// 		}
// 	}
// 	return words
// }

func main() {
	value1, _ := hexToDecimal("1E")
	value2, _ := hexToDecimal("FF")
	fmt.Println(value1)
	fmt.Println(value2)
	fmt.Println(binToDecimal("10"))
	fmt.Println(binToDecimal("1010"))
	fmt.Println(binToDecimal("11111111"))
	fmt.Println(capitalize("hELLO"))
	fmt.Println(capitalize("wORLD"))
	fmt.Println(UpperCaseLastN([]string{"This", "is", "so", "exciting"}, 2))
	fmt.Println(isPunc("!"))
	fmt.Println(isPunc(","))
	fmt.Println(isPunc("x"))
	fmt.Println(fixPunctuation("Hello , World !"))
	fmt.Println(fixArticles("a hour a fool A antelope"))
	fmt.Println(aOrAn("apple"))
	fmt.Println(aOrAn("hour"))
	fmt.Println(aOrAn("ball"))
	fmt.Println(fixSingleQuotes("'        Hello      World         '    "))
	fmt.Println(fixSingleQuotes("'           Awesome      ' "))

}

// import (
// 	"fmt"
// 	"strings"
// )

// func capEven(s string) string {
// 	if len(s)%2 == 0 {
// 		return strings.ToUpper(s)
// 	}
// 	return s
// }
// func main() {
// 	fmt.Println(capEven("Hello"))
// 	fmt.Println(capEven("Hell"))
// }
