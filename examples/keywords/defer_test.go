package keywords

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCounter_BadAdd(t *testing.T) {
	c := NewCounter()
	ret := 0
	for {
		tmp := c.BadAdd()
		if ret == tmp {
			break
		}
		ret = tmp
	}
	assert.Equal(t, 1001, ret)
}

func TestCounter_GoodAdd(t *testing.T) {
	c := NewCounter()
	ret := 0
	for {
		tmp := c.BadAdd()
		if ret == tmp {
			break
		}
		ret = tmp
	}
	assert.Equal(t, 1001, ret)
}

func BenchmarkCounter_BadAdd(b *testing.B) {
	c := NewCounter()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.BadAdd()
		}
	})
}

func BenchmarkCounter_GoodAdd(b *testing.B) {
	c := NewCounter()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.GoodAdd()
		}
	})
}
