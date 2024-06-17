package main

import "fmt"

// удаление с сохранением порядка и сохранением массива
func RemoveElementOrder(m []int, i int) ([]int, error) {

	if i < 0 || i >= len(m) {
		return nil, fmt.Errorf("индекс вне диапазона")
	}

	// создаем новый слайс чтобы не менять изначальный массив
	new := []int{}
	new = append(new, m[:i]...)
	new = append(new, m[i+1:]...)
	return new, nil

}

// удаление без сохранения порядка и без сохранения массива
func RemoveElement(m []int, i int) ([]int, error) {

	if i < 0 || i >= len(m) {
		return nil, fmt.Errorf("индекс вне диапазона")
	}

	m[i] = m[len(m)-1]
	return m[:len(m)-1], nil

}

func main() {

	x := []int{4, 5, 0, 9, 15, 81, -3}
	x2 := []int{4, 5, 0, 9, 15, 81, -3}

	y, _ := RemoveElementOrder(x, 3)
	y2, _ := RemoveElement(x2, 3)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(x2)
	fmt.Println(y2)

}
