package main

import (
	"fmt"
	"strconv"
	"sync"
)

// обернем мапу в структуру с мъютексом, чтобы лочить мапу при записи
type WSyncMap struct {
	m  map[string]string
	mu sync.RWMutex
}

func New() *WSyncMap {
	return &WSyncMap{
		m:  make(map[string]string),
		mu: sync.RWMutex{},
	}
}

// лочим мапу при записи в нее
func (wsm *WSyncMap) SafeWrite(key, value string) {
	wsm.mu.Lock()
	defer wsm.mu.Unlock()
	wsm.m[key] = value
}

func main() {

	wg := sync.WaitGroup{}

	syncMap := WSyncMap{m: make(map[string]string)}

	for i := 0; i <= 10; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()

			index := strconv.Itoa(i)
			syncMap.SafeWrite("key"+index, "value"+index)
		}()
	}

	wg.Wait()
	fmt.Printf("Результат: %v", syncMap.m)
}
