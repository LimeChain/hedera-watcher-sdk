package server

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/limechain/hedera-watcher-sdk/handlers"
	"github.com/limechain/hedera-watcher-sdk/interfaces"
	"github.com/limechain/hedera-watcher-sdk/queue"
	"github.com/limechain/hedera-watcher-sdk/types"
	"log"
	"net/http"
)

type HederaWatcherServer struct {
	handler  *handlers.Handler
	queue    *queue.Queue
	watchers []interfaces.Watcher
}

func (server *HederaWatcherServer) Run(addr string) {
	server.start()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	log.Println(fmt.Sprintf("Listening on port %s", addr))
	log.Fatal(http.ListenAndServe(addr, r))
}

func (server *HederaWatcherServer) AddWatcher(watcher interfaces.Watcher) {
	server.watchers = append(server.watchers, watcher)
}

func (server *HederaWatcherServer) AddHandler(messageType string, handler interfaces.Handler) {
	server.handler.AddHandler(messageType, handler)
}

func (server *HederaWatcherServer) start() {
	ch := make(chan *types.Message)
	server.queue = queue.NewQueue(ch)
	go server.handler.Handle(ch)
	go server.startWatchers()
}

func (server *HederaWatcherServer) startWatchers() {
	go func() {
		for _, watcher := range server.watchers {
			watcher.Watch(server.queue)
		}
	}()
}

func NewServer() *HederaWatcherServer {
	return &HederaWatcherServer{
		handler: handlers.NewHandler(),
	}
}
