package builder

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewPool(t *testing.T) {

	type Item struct {
		name    string
		opts    []OptionFunc
		expect  *Pool
		isError bool
	}

	items := []Item{
		{
			name:    "",
			expect:  nil,
			isError: true,
		}, {
			name: "test",
			opts: []OptionFunc{
				func(option *PoolOption) {
					option.minIdle = 2
				},
			},
			expect: &Pool{
				option: &PoolOption{
					maxTotal: 10,
					maxIdle:  9,
					minIdle:  2,
				},
			},
			isError: false,
		},
	}

	for _, item := range items {
		t.Run(item.name, func(t *testing.T) {
			pool, err := NewPool(item.name, item.opts...)
			require.Equalf(t, item.isError, err != nil,
				"error=%v, expect error=%v", err, item.isError)
			assert.Equal(t, item.expect, pool)
		})
	}
}
