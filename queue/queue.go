package queue

import "github.com/limechain/hedera-watcher-sdk/types"

type Queue struct {
	channel chan *types.Message
}

func (q *Queue) Push(message *types.Message) {
	q.channel <- message
}

func NewQueue(ch chan *types.Message) *Queue {
	return &Queue{channel: ch}
}
