package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value  int
	wgroup sync.WaitGroup
}

func (c *Counter) inc() {
	c.value = c.value + 1
	c.wgroup.Done()
}

func main() {
	// OBS: Executar em um ambiente com vários núcleos de processamento
	var c Counter
	for i := 0; i < 1000; i++ {
		c.wgroup.Add(1)
		go c.inc()
	}
	c.wgroup.Wait()
	fmt.Println(c.value)
}
