package queue

import (
	"testing"
)

type TestMessage struct {
	MessageData
	data string
}

func (m TestMessage) GetData() (any, error) {
	return m.data, nil
}

func TestTransformerQueue_Crud(t *testing.T) {
	queue := LocalQueue{
		queue: make(chan Message, 1),
	}

	message := TestMessage{
		MessageData: NewMessageData(),
		data:        "test",
	}
	queue.Add(message)
	<-queue.Chan()
	queue.Add(message)
	<-queue.Chan()
	queue.Close()
}
