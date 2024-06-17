package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	defer fmt.Println("Программа остановлена")

	// создаем контекст с отменой по таймауту
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// освобождаем ресурсы если программа завершится до таймаута, вызовет ctx.Done.
	// в данной реальзации не актуально т.к. данные в канал пишутся в бесконечном цикле.
	defer cancel()

	// синхронизация необходима чтобы дождаться окончания работы
	// горутины которая читает данные из канала и после завешить программу
	wg := sync.WaitGroup{}

	// канал для данных
	data := make(chan int)

	// горутина которая читает данные из канала
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		for {
			select {
			case d := <-data:
				fmt.Println(d)
			case <-ctx.Done():
				return
			}
		}
	}(&wg)

	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			wg.Wait()
			return
		case data <- rand.Intn(999):
			time.Sleep(1 * time.Second)
		}
	}
}
