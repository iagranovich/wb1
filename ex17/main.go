package main

import "fmt"

// возвращаем индекс входа в массив, -1 если не найдено
func binarysearch(n int, x []int) int {
	low := 0
	high := len(x) - 1

	for low <= high {

		mid := low + (high-low)/2

		if x[mid] == n {
			return mid
		} else if x[mid] < n {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1

}

func main() {
	x := []int{-1, 4, 5, 9, 11}

	fmt.Println(binarysearch(-1, x))
	fmt.Println(binarysearch(0, x))
	fmt.Println(binarysearch(11, x))
}
