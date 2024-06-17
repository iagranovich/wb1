package main

import "fmt"

type SDCard struct {
}

func NewSD() *SDCard {
	return &SDCard{}
}

func (sd *SDCard) ConnectToSD() {
	fmt.Println("Карта подключена")
}

type USB interface {
	ConnectToUSB()
}

func Connection(u USB) {
	u.ConnectToUSB()
	fmt.Println("Устройство USB подключено")
}

func main() {
	sd := NewSD()
	adapter := NewAdapter(sd)

	Connection(adapter)
}

type USBAdapter struct {
	sd *SDCard
}

func NewAdapter(sd *SDCard) *USBAdapter {
	return &USBAdapter{sd: sd}
}

// реализация USB interface
func (a *USBAdapter) ConnectToUSB() {
	a.sd.ConnectToSD()
}
