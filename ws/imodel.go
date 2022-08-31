package ws

type IModel interface {
	Emit(event string, msg interface{}) error
	Disconnect(reason string) error
	Ping() error
}