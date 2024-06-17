package main

import "fmt"

func main() {
	set1 := []int{0, 2, 3, 1}
	set2 := []int{4, 3, 2, 5, 6}

	// создадим вспомагательную мапу, у которой
	// ключами будут элементы первого множества
	// а все значения bool для удобства последующей проверки
	helper := make(map[int]bool)

	for _, num := range set1 {
		helper[num] = true
	}

	// проверим есть ли среди ключей элементы второго множества
	intersection := []int{}
	for _, num := range set2 {
		if helper[num] {
			intersection = append(intersection, num)
		}
	}

	fmt.Println(intersection)
}
