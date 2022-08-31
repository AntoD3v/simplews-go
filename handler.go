package simplews_go

import (
	"github.com/antod3v/simplews_go/ws"
)


type HandleType int
const (
	handleConnect    HandleType = 0
	handleDisconnect HandleType = 1
	handleError      HandleType = 2
	handleEvent      HandleType = 3
)

type Handles struct {
	handleConnect []func(customer ws.Customer) error
	handleDisconnect []func(customer ws.Customer, reason string)
	handleError []func(err error)
	handleEvent map[string][]func(customer * ws.Customer, msg interface{})
}

func (h *Server) On(event string, f func(customer * ws.Customer, msg interface{})) {
	h.OnEvent(event, f)
}

func (h *Server) OnEvent(event string, f func(customer * ws.Customer, msg interface{})) {
	h.handles.handleEvent[event] = append(h.handles.handleEvent[event], f)
}

func (h *Server) OnConnect(f func(customer ws.Customer) error) {
	h.handles.handleConnect = append(h.handles.handleConnect, f)
}

func (h *Server) OnDisconnect(f func(customer ws.Customer, reason string)) {
	h.handles.handleDisconnect = append(h.handles.handleDisconnect, f)
}

func (h *Server) OnError(f func(err error)) {
	h.handles.handleError = append(h.handles.handleError, f)

}

func (h *Server) call(handleType HandleType, args ...interface{}) {

	switch handleType {
	case handleConnect:
		for _, f := range h.handles.handleConnect {
			f(args[0].(ws.Customer))
		}
	case handleDisconnect:
		for _, f := range h.handles.handleDisconnect {
			f(args[0].(ws.Customer), args[1].(string))
		}
	case handleError:
		for _, f := range h.handles.handleError {
			f(args[0].(error))
		}
	case handleEvent:
		for _, f := range h.handles.handleEvent[args[0].(string)] {
			f(args[1].(*ws.Customer), args[2].(interface{}))
		}

	}
}
