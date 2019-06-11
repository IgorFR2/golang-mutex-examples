package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	value  int
	wgroup sync.WaitGroup
}

func (c *Counter) inc() {
	time.Sleep(1 * time.Second)
	c.value = c.value + 1
	c.wgroup.Done()
}

func main() {
	var c Counter
	for i := 0; i < 1000; i++ {
		c.wgroup.Add(1)
		go c.inc()
	}
	c.wgroup.Wait()
	fmt.Println(c.value)
}
