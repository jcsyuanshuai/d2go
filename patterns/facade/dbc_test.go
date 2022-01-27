package facade

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	api := NewDbc()
	ret := api.Connect()
	assert.Equal(t, "connection of mysql.\nconnection of mongo.", ret)
}
