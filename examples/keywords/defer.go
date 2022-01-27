package keywords

import "sync"

type Counter struct {
	mu      sync.Mutex
	counter int
}

func NewCounter() *Counter {
	return &Counter{
		counter: 0,
	}
}

func (c *Counter) BadAdd() int {
	c.mu.Lock()

	if c.counter > 1000 {
		c.mu.Unlock()
		return c.counter
	}

	c.counter++
	ret := c.counter
	c.mu.Unlock()
	return ret
}

func (c *Counter) GoodAdd() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.counter > 1000 {
		return c.counter
	}
	c.counter++
	return c.counter
}
