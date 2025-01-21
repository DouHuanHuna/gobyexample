package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()
	c.value++
}

func (c *Counter) Get() int {
	c.Lock()
	defer c.Unlock()
	return c.value
}

func main() {
	counter := &Counter{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()

	}

	wg.Wait()

	fmt.Println(counter.Get())
}
