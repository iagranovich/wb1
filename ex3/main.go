package main

import (
	"sync"
)

var numbers = []int{2, 4, 6, 8, 10}

func v1() {
	mx := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(numbers))
	sum := 0

	// на каждое число запускаем горутину
	for _, num := range numbers {
		num := num
		go func() {
			defer wg.Done()

			// используем мъютекс чтоьы не было гинки
			mx.Lock()
			defer mx.Unlock()
			sum += num * num

		}()
	}
	wg.Wait()

	//fmt.Println(sum)

}

func v2() {
	sqrs := make(chan int)

	// считаем квадраты в горутине и записываем в канал
	go func() {
		for _, num := range numbers {
			sqrs <- num * num
		}
		// закрываем канал когда числа кончились чтобы не было дедблока
		close(sqrs)
	}()

	// читаем из канала и суммируем
	sum := 0
	for sqr := range sqrs {
		sum += sqr
	}

	//fmt.Println(sum)

}

func v3() {
	// создаем буферизированный канал
	sqrs := make(chan int, len(numbers))

	// считаем квадраты в горутине и записываем в канал
	go func() {
		for _, num := range numbers {
			sqrs <- num * num
		}
		// закрываем канал когда числа кончились чтобы не было дедблока
		close(sqrs)
	}()

	// читаем из канала и суммируем
	sum := 0
	for sqr := range sqrs {
		sum += sqr
	}

	//fmt.Println(sum)

}

func main() {
	v1()
	v2()
	v3()
}
