package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "aogbanje", Age: 35}
	// p := Person{}
	// p.Name = "aogbanje"
	// p.Age = 35
	fmt.Println(p.Name)
	fmt.Println(p.Age)
}
