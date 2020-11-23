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

func (t *WatchHandler) Handle(payload []byte) {
	log.Println("handling message...", payload)
}

func (t *WatchHandler) Watch(queue *queue.Queue) {
	go func() {
		for true {
			msg := &types.Message{
				Payload: []byte("[1] pushing message"),
				Type:    t.messageType,
			}
			log.Println("[1] pushing message...")
			queue.Push(msg)
			time.Sleep(1 * time.Second)
		}
	}()
}

func (t *WatchHandler) Recover(queue *queue.Queue) {
	log.Println("recovering...")
	queue.Push(&types.Message{
		Payload: []byte("recovery message"),
		Type:    t.messageType,
	})
}

func NewWatcherHandler(messageType string) *WatchHandler {
	return &WatchHandler{messageType}
}
