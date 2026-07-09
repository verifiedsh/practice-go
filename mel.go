package main

import (
	"fmt"
	//"unicode"
	"strings"
	//"bufio"
	//"os"
)

func Capit() {
	name := ""
	fmt.Print("Enter Name: ")
	fmt.Scan(&name)
	fmt.Println("Hello " + strings.ToUpper(name))
	// return "Hello" + name
}
