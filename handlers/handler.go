package handlers

import (
	"github.com/limechain/hedera-watcher-sdk/interfaces"
	"github.com/limechain/hedera-watcher-sdk/types"
	"log"
)

type Handler struct {
	handlers map[string]interfaces.Handler
}

func (h *Handler) Handle(ch <-chan *types.Message) {
	for message := range ch {
		if err := h.handleMessage(message); err != nil {
			log.Println(err.Error())
		}
	}
}

func (h *Handler) handleMessage(msg *types.Message) error {
	handler := h.handlers[msg.Type]
	if err := handler.Handle(msg.Payload); err != nil {
		return err
	}

	return nil
}

func (h *Handler) AddHandler(messageType string, handler interfaces.Handler) {
	h.handlers[messageType] = handler
}

func NewHandler() *Handler {
	handlers := make(map[string]interfaces.Handler)
	return &Handler{handlers}
}
