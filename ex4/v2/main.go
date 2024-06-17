package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Аналогично, первому варианту, но без контекста. При получении сигнала о завершении
// работы программы, закрываем канал с данными. Это приводит к завершению работы воркеров.
// После этого завершается работа программы.
func worker(id int, data chan int, wg *sync.WaitGroup) {
	defer func() {
		fmt.Printf("Воркер %d остановлен\n", id)
		wg.Done()
	}()

	for d := range data {
		fmt.Printf("Воркер %d: %d\n", id, d)
	}
}

func main() {
	wg := sync.WaitGroup{}

	data := make(chan int)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT)

	fmt.Print("Введите необходимое количество воркеров: ")
	var N int
	fmt.Scan(&N)

	for n := 1; n <= N; n++ {
		n := n
		wg.Add(1)
		go func() {
			worker(n, data, &wg)
		}()
	}

	for {
		select {
		case <-interrupt:
			close(data)
			wg.Wait()
			fmt.Println("Приложение завершается")
			return

		case data <- rand.Intn(999):
			time.Sleep(time.Second)
		}
	}

}
