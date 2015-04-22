package main

import (
	"testing"
)

func TestMessageID(t *testing.T) {
	b := NewBroker()
	m := NewMessage(b.NewID(), 1)
	if len(m.ID) <= 0 {
		t.Errorf("messageID is %v", m.ID)
	}

	t.Logf("%s", m.ID)
}
