package main

import (
	"context"
	"fmt"
	"time"
)

func StopByTimeout(ctx context.Context) {
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

	// создаем контекст с таймаутом, по истечении которого
	// будет послан сигнал об остановке в канал ctx.Done
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	// запускаем горутину
	go StopByTimeout(ctx)

	// чтобы увидеть эффект зададим таймаут больше чем в контексте
	time.Sleep(time.Second * 5)
}
