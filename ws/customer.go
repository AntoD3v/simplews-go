package ws

import (
	"github.com/antod3v/simplews_go/parser"
	"github.com/gorilla/websocket"
)

type Customer struct {
	encoder *parser.Encoder
	conn *websocket.Conn
}

func CreateCustomer(encoder * parser.Encoder, conn * websocket.Conn) Customer {
	return Customer{encoder, conn}
}

// Implement IModel with Customer
func (c *Customer) Emit(event string, msg interface{}) error {
	model := parser.Model{
		Type: parser.Event,
		Data: &msg,
		Event: &event,
	}
	data, err := c.encoder.Encode(&model)
	if err != nil {
		return err
	}
	return c.conn.WriteMessage(websocket.TextMessage, data)
}

func (c *Customer) Disconnect(reason string) error {
	model := parser.Model{
		Type: parser.Disconnect,
		Message: &reason,
	}
	data, err := c.encoder.Encode(&model)
	if err != nil {
		return err
	}
	return c.conn.WriteMessage(websocket.CloseMessage, data)
}

func (c *Customer) Ping() error {
	return nil
}
