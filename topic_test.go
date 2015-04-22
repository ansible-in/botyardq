package main

import (
	"sync"
	"testing"
)

func TestTopicSimple(t *testing.T) {
	topic := NewTopic("simple")
	var id MessageID
	m := NewMessage(id, 1)
	topic.PushMessage(m)
	data := topic.PopMessage()
	if data == nil {
		t.Errorf("data is nill,wants 1")
	}

	if data.Value.(int) != 1 {
		t.Errorf("data is not 1")
	}
}

func _TestTopicPopBlock(t *testing.T) {
	var wg sync.WaitGroup

	topic := NewTopic("simple")
	go func() {
		var id MessageID
		m := NewMessage(id, 1)
		topic.PushMessage(m)
	}()

	go func() {
		wg.Add(1)
		m := topic.PopMessage()
		if m.Value.(int) != 1 {
			t.Errorf("PopMessage should be blocked %v", m)
		}
		wg.Done()
	}()
	wg.Wait()

}

func BenchmarkTopicPush(b *testing.B) {
	topic := NewTopic("simple")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var id MessageID
		m := NewMessage(id, i)
		topic.PushMessage(m)
	}
}
