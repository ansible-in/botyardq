package main

import (
	"log"
	"sync"
)

type Topic struct {
	Name      string
	Queue     Queue
	waitingCh chan interface{}
	mutex     sync.Mutex
}

func NewTopic(name string) *Topic {

	topic := &Topic{
		Name:      name,
		Queue:     NewLLQueue(),
		waitingCh: make(chan interface{}),
	}

	return topic
}

func (t *Topic) PushMessage(msg interface{}) {
	select {
	case t.waitingCh <- msg:
	default:
		t.mutex.Lock()
		defer t.mutex.Unlock()
		err := t.Queue.Push(msg)
		if err != nil {
			//TODO:
			log.Println(err)
		}
	}
}

func (t *Topic) PopMessage() interface{} {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if msg := t.Queue.Pop(); msg != nil {
		return msg
	} else {
		msg := <-t.waitingCh
		return msg
	}
}
