package main

import (
	"fmt"
	"syscall"
	"time"
)

// используем пакет time
func TimeSleep(duration time.Duration) {
	<-time.After(duration)
}

// используем пакет syscall. используемый системный вызов nanosleep
// поддерживается только операционными системами стандарта POSIX
func SyscallSleep(seconds int) {

	ts := syscall.Timespec{
		Sec:  int64(seconds),
		Nsec: 0,
	}

	err := syscall.Nanosleep(&ts, nil)
	if err != nil {
		fmt.Println("ошибка при вызове nanosleep:", err)
	}
}

func main() {
	fmt.Println("Пора спать.")
	TimeSleep(2 * time.Second)
	fmt.Println("Доброе утро!")
	SyscallSleep(1)
	fmt.Println("Добрый день!")
}
