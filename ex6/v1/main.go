package main

import (
	"fmt"
	"time"
)

func StopByCloseChannel(out chan int) {
	for range out {
		fmt.Println("Горутина что-то делает ...")
	}
	fmt.Println("Горутина остановлена")
}

func main() {
	// канал который горутина читает и который мы закроем
	data := make(chan int)

	// запускаем горутину
	go StopByCloseChannel(data)

	// если закроем канал, то цикл for range на чтение
	// из канала завершится и горутина завершится
	time.Sleep(time.Second * 1)
	close(data)
	time.Sleep(time.Second * 1)
}
