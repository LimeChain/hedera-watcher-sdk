package interfaces

type Handler interface {
	Handle([]byte)
}
