package interfaces

import "github.com/limechain/hedera-watcher-sdk/queue"

type Handler interface {
	Handle([]byte)
	Recover(queue *queue.Queue)
}
