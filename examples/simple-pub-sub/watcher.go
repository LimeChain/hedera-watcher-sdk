package main

import (
	"github.com/limechain/hedera-watcher-sdk/queue"
	"github.com/limechain/hedera-watcher-sdk/types"
	"log"
	"time"
)

type Watcher struct {
	messageType string
}

func (t *Watcher) Watch(queue *queue.Queue) {
	go func() {
		for true {
			msg := &types.Message{
				Payload: nil,
				Type:    t.messageType,
			}
			log.Println("[2] pushing message...")
			queue.Push(msg)
			time.Sleep(1 * time.Second)
		}
	}()
}

func NewWatcher(messageType string) *Watcher {
	return &Watcher{messageType}
}
