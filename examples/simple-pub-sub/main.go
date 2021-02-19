package main

import (
	"github.com/go-chi/chi"
	"github.com/limechain/hedera-watcher-sdk/server"
)

const (
	MessageType = "PUB_SUB"
)

func main() {
	router := chi.NewRouter()
	wh := NewWatcherHandler(MessageType)
	w := NewWatcher(MessageType)

	watcherServer := server.NewServer()
	watcherServer.AddHandler(wh.messageType, wh)
	watcherServer.AddWatcher(wh)
	watcherServer.AddWatcher(w)

	watcherServer.Run(router, ":3000")
}
