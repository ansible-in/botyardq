package main

import (
	"log"
	"runtime"
	"sync"
	"time"
)

type Broker struct {
	ID         int64
	Topics     map[string]*Topic
	topicMutex sync.Mutex
	idChan     chan MessageID
}

func NewBroker() *Broker {

	b := &Broker{
		ID:     1,
		idChan: make(chan MessageID, 4096), // Buffer
	}

	go b.idPump()

	return b
}

func (b *Broker) Topic(name string) *Topic {
	b.topicMutex.Lock()
	defer b.topicMutex.Unlock()

	if t, ok := b.Topics[name]; ok {
		return t
	}

	t := NewTopic(name)
	b.Topics[name] = t

	return t
}

func (b *Broker) PushMessage(name string, value interface{}) error {
	t := b.Topic(name)
	msg := NewMessage(b.NewID(), value)
	t.PushMessage(msg)
	return nil
}

func (b *Broker) PopMessage(name string) *Message {
	t := b.Topic(name)
	msg := t.PopMessage()
	return msg
}

//This method is implemented to MessageIDGenerator
func (b *Broker) NewID() MessageID {
	return <-b.idChan
}

func (b *Broker) idPump() {
	factory := &guidFactory{}
	lastError := time.Now()

	for {

		id, err := factory.NewGUID(b.ID)

		if err != nil {
			now := time.Now()
			if now.Sub(lastError) > time.Second {
				//only print the error once/second
				//TODO:
				lastError = now
				log.Printf("ERROR: %s", err)

			}
			runtime.Gosched()
			continue
		}

		select {
		case b.idChan <- id.Hex():
			//TODO: exitChan?
		}
	}
}
