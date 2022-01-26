package decorator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestColoredSquare_Draw(t *testing.T) {
	s := NewColoredSquare(&Square{}, "red")
	ret := s.Draw()
	assert.Equal(t, "this is a square, color is red", ret)
}
