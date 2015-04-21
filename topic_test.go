package main

import (
	"sync"
	"testing"
)

func TestTopicSimple(t *testing.T) {
	topic := NewTopic("simple")
	topic.PushMessage(1)
	data := topic.PopMessage()
	if data == nil {
		t.Errorf("data is nill,wants 1")
	}

	if data.(int) != 1 {
		t.Errorf("data is not 1")
	}
}

func _TestTopicPopBlock(t *testing.T) {
	var wg sync.WaitGroup

	topic := NewTopic("simple")
	go func() {
		topic.PushMessage(1)
	}()

	go func() {
		wg.Add(1)
		data := topic.PopMessage()
		if data.(int) != 1 {
			t.Errorf("PopMessage should be blocked %v", data)
		}
		wg.Done()
	}()
	wg.Wait()

}

func BenchmarkTopicPush(b *testing.B) {
	topic := NewTopic("simple")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		topic.PushMessage(i)
	}
}
