package main

import (
	"context"
	"fmt"
	"time"
)

func StopBySignalFromContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Горутина остановлена")
			return
		default:
			fmt.Println("Горутина что-то делает...")
			time.Sleep(time.Second)
		}
	}
}

func main() {

	// создаем контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())

	// запускаем горутину
	go StopBySignalFromContext(ctx)

	// вызываем специальную функцию отмены контекста
	// которая отправит в канал ctx.Done сигнал остановки
	time.Sleep(time.Second * 1)
	cancel()
	time.Sleep(time.Second * 2)
}
