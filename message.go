package main

import (
	"time"
)

const MsgIdLen = 16

const (
	MSG_ENQUEUED = iota
	MSG_DEQUEUED
	MSG_FINISHED
)

type MessageID [MsgIdLen]byte

type MessageIDGenerator interface {
	NewID() MessageID
}

type Message struct {
	ID      MessageID
	Value   interface{}
	Created time.Time
	State   uint8
}

func NewMessage(id MessageID, value interface{}) *Message {
	m := &Message{
		ID:      id,
		Value:   value,
		Created: time.Now(),
		State:   MSG_ENQUEUED,
	}

	return m
}
