package bridge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWarnNotification_Notify(t *testing.T) {
	sender := NewEmailSender([]string{"test-1@test.com"})
	n := NewWarnNotification(sender)
	err := n.Notify("msg ...")
	assert.Nil(t, err)
}

func TestFailNotification_Notify(t *testing.T) {
	sender := NewWechatSender([]string{"Bob"})
	n := NewFailNotification(sender)
	err := n.Notify("msg ...")
	assert.Nil(t, err)
}
