package simplews_go

import (
	"log"
	"net/http"

	"github.com/antod3v/simplews_go/parser"
	"github.com/antod3v/simplews_go/ws"
	"github.com/gorilla/websocket"
)

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var upgrader = websocket.Upgrader{
		CheckOrigin:     *s.opts.CheckOrigin,
		ReadBufferSize:  *s.opts.ReadBufferSize,
		WriteBufferSize: *s.opts.WriteBufferSize,
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	defer c.Close()

	customer := ws.CreateCustomer(s.encoder, c)
	s.call(handleConnect, customer)

	for {

		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("Unable to read message", err)
			break
		}

		model, err := s.decoder.Decode(message)
		if err != nil {
			log.Println("Unable to decode message", err)
			break
		}

		switch mt {

		case websocket.TextMessage:
			switch model.Type {

			case parser.Event:
				s.call(handleEvent, *model.Event, &customer, model.Data)
				break

			case parser.Disconnect:
				s.call(handleDisconnect)
				break

			case parser.Error:
				s.call(handleError)
				break
			}

		case websocket.CloseMessage:
			s.call(handleDisconnect, message)

		default:
			s.call(handleDisconnect, "Unknown message close")

		}
	}

}
