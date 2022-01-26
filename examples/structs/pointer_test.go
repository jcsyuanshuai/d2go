package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointerToInterface(t *testing.T) {
	var s1 = NewS1(1) // struct 值传递
	var s2 = NewS2(1) // struct 指针传递

	var m1 Modifier = s1
	var m2 Modifier = s2

	m1.change()
	m2.change()

	assert.NotEqual(t, s1.Value, s2.Value)
}
