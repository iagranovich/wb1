package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Данный воркер исполузует контекст для отслеживания сигнала о
// завершении программы, что позволяет освободить ресурсы, связанные
// с данным контекстом. Таким образом, получив сигнал о прекращении работы,
// мы в начале завершаем работу воркеров, а затем работу самой программы.
func worker(ctx context.Context, id int, data chan int, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		fmt.Printf("Воркер %d остановлен\n", id)
	}()

	// читаем из канала данные, пока в канал ctx.Done
	// не придет сигнал о завершении работы
	for {
		select {
		case <-ctx.Done():
			return
		case d := <-data:
			fmt.Printf("Воркер %d: %d\n", id, d)
		}
	}
}

func main() {
	defer fmt.Println("Программа остановлена")

	// синхронизация необходима чтобы дождаться окончания работы
	// всех горутин, прежде чем закончить работу пограммы
	wg := sync.WaitGroup{}

	data := make(chan int)

	// контекст
	ctx, cancel := context.WithCancel(context.Background())
	// канал для сигнала окончания работы программы
	interrupt := make(chan os.Signal, 1)
	// передача сигнала в канал
	signal.Notify(interrupt, syscall.SIGINT)

	// в отдельной горутине отслежтваем появление сигнала на завершение
	go func() {
		<-interrupt
		// отменяем контекст и отправляем сигнал в канал ctx.Done
		cancel()
	}()

	// обрабатывем ввод с консоли
	fmt.Print("Введите необходимое количество воркеров: ")
	var N int
	fmt.Scan(&N)

	// запускаем необходимое количество воркеров
	for n := 1; n <= N; n++ {
		n := n
		wg.Add(1)
		go func() {
			worker(ctx, n, data, &wg)
		}()
	}

	// отправляем данные в канал пока не получим сигнал о завершении из канла ctx.Done
	for {
		select {
		case <-ctx.Done():
			// ждем пока все воркеры закончат работу
			wg.Wait()
			// завершаем программу
			return

		case data <- rand.Intn(999):
			time.Sleep(time.Second)
		}
	}

}
