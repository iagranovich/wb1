package main

import "fmt"

func main() {
	numbers := [10]int{0, 1, 2, 3, -4, 5, 6, 7, 8, -11}

	// создаем каналы для конвеера
	nums := make(chan int, len(numbers))
	doubles := make(chan int, len(numbers))

	// передаем числа в канал для чисел
	go func() {
		for _, num := range numbers {
			nums <- num
		}
		// закрываем канал когда числа кончились чтобы не было дедблока
		close(nums)
	}()

	// читаем из канала чисел, умножаем и передаем в следующий канал
	go func() {
		for num := range nums {
			doubles <- num * 2
		}
		// закрываем канал когда числа кончились чтобы не было дедблока
		close(doubles)
	}()

	// читаем из канала и печатаем
	for n := range doubles {
		fmt.Println(n)
	}

}
