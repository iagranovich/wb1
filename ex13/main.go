package main

import "fmt"

func main() {
	a := -1
	b := 1

	b = a - b
	a = a - b
	b = a + b

	fmt.Println(a, b)
}
