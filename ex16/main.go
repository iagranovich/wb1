package main

import "fmt"

func quicksort(x []int) []int {

	if len(x) <= 1 {
		return x
	}

	p := partition(x)

	q1 := quicksort(x[:p])
	q2 := quicksort(x[p:])

	return append(q1, q2...)
}

func partition(x []int) int {

	last := len(x) - 1
	pivot := x[last]

	j := 0
	for i := 0; i < last; i++ {
		if x[i] <= pivot {

			x[i], x[j] = x[j], x[i]
			j++
		}
	}
	x[last], x[j] = x[j], x[last]

	return j
}

func main() {

	x := []int{4, -2, 25, 1, 0, 7}

	x = quicksort(x)

	fmt.Println(x)

}
