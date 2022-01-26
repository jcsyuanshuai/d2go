package keywords

import "sync"

type Counter struct {
	mu      sync.Mutex
	counter int
}

func badAdd() int {
	c := Counter{}
	c.mu.Lock()

	if c.counter > 10 {
		c.mu.Unlock()
		return c.counter
	}

	c.counter++
	ret := c.counter
	c.mu.Unlock()
	return ret
}

func goodAdd() int {
	c := Counter{}
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.counter > 10 {
		return c.counter
	}
	c.counter++
	return c.counter
}
