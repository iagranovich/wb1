package main

import (
	"fmt"
	"sync"
)

// используем мъютекс для конкурентного доступа к счетчику
type Counter struct {
	i  int
	mx sync.Mutex
}

func New() *Counter {
	return &Counter{}
}

func (c *Counter) Count() {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.i++
}

func main() {
	// для синхронизации горутин
	wg := sync.WaitGroup{}

	counter := New()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Count()
		}()
	}
	wg.Wait()
	fmt.Println(counter.i)
}
