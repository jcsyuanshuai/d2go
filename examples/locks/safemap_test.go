package locks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSafeMap(t *testing.T) {
	smp := NewSafeMap()

	smp.Put("test-key", "test-val")

	ret := smp.Get("test-key")

	assert.Equal(t, "test-val", ret)

}
