package main

import "fmt"

func main() {
	// последовательность
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	// создадим вспомагательную мапу, в которой
	// ключи будут элементами последовательности
	helper := make(map[string]struct{})

	for _, s := range sequence {
		helper[s] = struct{}{}
	}

	// сохраняя ключи в слайс, получим множество
	// для заданной последовательности
	set := []string{}
	for k := range helper {
		set = append(set, k)
	}

	fmt.Println(set)
}
