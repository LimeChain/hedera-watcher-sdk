package main

import (
	"github.com/limechain/hedera-watcher-sdk/queue"
	"github.com/limechain/hedera-watcher-sdk/types"
	"log"
	"time"
)

type WatchHandler struct {
	messageType string
}

func (t *WatchHandler) Handle([]byte) error {
	log.Println("handling message...")
	return nil
}

func (t *WatchHandler) Watch(queue *queue.Queue) {
	go func() {
		for true {
			msg := &types.Message{
				Payload: nil,
				Type:    t.messageType,
			}
			log.Println("[1] pushing message...")
			queue.Push(msg)
			time.Sleep(1 * time.Second)
		}
	}()
}

func NewWatcherHandler(messageType string) *WatchHandler {
	return &WatchHandler{messageType}
}
