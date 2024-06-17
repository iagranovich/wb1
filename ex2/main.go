package main

import (
	"fmt"
	"sync"
)

func sq() {
	numbers := []int{2, 4, 6, 8, 10}
	// для синхронизации горутин
	wg := sync.WaitGroup{}
	wg.Add(len(numbers))

	// каждое число возводим в квадрат и выводим на консоль в отдельной горутине
	for _, num := range numbers {
		num := num
		go func() {
			defer wg.Done()
			fmt.Println(num * num)
		}()
	}

	// ждем когда горутины закончат выполнение
	wg.Wait()

}

func main() {
	sq()
}
