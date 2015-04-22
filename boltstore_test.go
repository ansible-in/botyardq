package main

import (
	"os"
	"testing"
)

func TestBoltStorePut(t *testing.T) {

	s, err := NewTopicStore("bolt", "/tmp", "test")
	if err != nil {
		t.Error(err)
	}
	err = s.Open()
	if err != nil {
		t.Error(err)
	}

	defer func() {
		s.Close()
		os.Remove(s.(*BoltStore).Path)
	}()

	b := NewBroker()
	m := NewMessage(b.NewID(), 1)

	err = s.PutMessage(m)
	if err != nil {
		t.Error(err)
	}

	var m1 *Message
	m1, err = s.GetMessage(m.ID)

	if m1.ID != m.ID {
		t.Errorf("ID %s,wants %s", m1.ID, m.ID)
	}

}
