package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	//s := longestPalindrome("babad")
	//
	//assert.Equal(t, "bab", s)

	s1 := longestPalindrome("bb")

	assert.Equal(t, "bb", s1)
}
