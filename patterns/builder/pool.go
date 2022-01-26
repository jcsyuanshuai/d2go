package builder

import (
	"errors"
	"fmt"
)

type PoolOption struct {
	maxTotal int
	maxIdle  int
	minIdle  int
}

type OptionFunc func(option *PoolOption)

type Pool struct {
	option *PoolOption
}

func NewPool(name string, opts ...OptionFunc) (*Pool, error) {
	if name == "" {
		return nil, errors.New("name cat not be empty")
	}

	option := &PoolOption{
		maxTotal: 10,
		maxIdle:  9,
		minIdle:  1,
	}

	for _, opt := range opts {
		opt(option)
	}

	wrongOption := errors.New(fmt.Sprintf("wrong option: %v", option))

	if option.maxTotal < 0 || option.maxIdle < 0 {
		return nil, wrongOption
	}
	if option.maxTotal < option.maxIdle || option.minIdle > option.maxIdle {
		return nil, wrongOption
	}

	return &Pool{
		option: option,
	}, nil
}
