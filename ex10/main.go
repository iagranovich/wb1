package main

import "fmt"

func main() {
	temps := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	subset := make(map[int][]float32)

	for _, temp := range temps {
		// ключ группы к которой пренадлежит температура = делим
		// температуру на 10, с отбрасыванием дробной части, и умножаем на 10
		key := int(temp/10) * 10

		subset[key] = append(subset[key], temp)
	}

	fmt.Println(subset)
}
