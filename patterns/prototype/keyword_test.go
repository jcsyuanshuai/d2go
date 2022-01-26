package prototype

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestKeywords_Clone(t *testing.T) {
	updatedAt, _ := time.Parse("2006", "2020")

	keywords := Keywords{
		"test-1": &Keyword{
			value:     "test-1",
			visit:     1,
			UpdatedAt: updatedAt,
		},
		"test-2": &Keyword{
			value:     "test-2",
			visit:     2,
			UpdatedAt: updatedAt,
		},
		"test-3": &Keyword{
			value:     "test-3",
			visit:     3,
			UpdatedAt: updatedAt,
		},
	}

	now := time.Now()

	updated := []*Keyword{
		{
			value:     "test-2",
			visit:     10,
			UpdatedAt: now,
		},
	}

	ret := keywords.Clone(updated)

	assert.Equal(t, keywords["test-1"], ret["test-1"])
	assert.Equal(t, keywords["test-3"], ret["test-3"])
	assert.NotEqual(t, keywords["test-2"], ret["test-2"])
	assert.NotEqual(t, updated[0], ret["test-2"])
}
