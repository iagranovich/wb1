package main

import (
	"fmt"
	"time"
)

func StopBySendingSignalInChannel(stop chan interface{}) {
	for {
		select {
		case <-stop:
			fmt.Println("Горутина остановлена")
			return
		default:
			fmt.Println("Горутина что-то делает...")
			time.Sleep(time.Second)
		}
	}
}

func main() {

	// специальный канал для сигнала остановки
	stop := make(chan interface{})

	// запускаем горутину
	go StopBySendingSignalInChannel(stop)

	// отправляем сигнал в канал
	time.Sleep(time.Second * 1)
	stop <- struct{}{}
	time.Sleep(time.Second * 1)
}
