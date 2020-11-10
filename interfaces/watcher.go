package interfaces

import (
	"github.com/limechain/hedera-watcher-sdk/queue"
)

type Watcher interface {
	Watch(queue *queue.Queue)
}
