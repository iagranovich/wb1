package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	i int32
}

func New() *Counter {
	return &Counter{}
}

// используем пакет sync/atomic
func (c *Counter) Count() {
	atomic.AddInt32(&c.i, 1)
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
