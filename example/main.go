package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/antod3v/simplews_go"
	"github.com/antod3v/simplews_go/ws"
)

func main() {
	
	simplews := simplews_go.New(&simplews_go.Opts{
		ReadBufferSize:  simplews_go.PtrInt(1024),
		WriteBufferSize: simplews_go.PtrInt(1024),
		Compression:     simplews_go.PtrBool(false),
		Base64:          simplews_go.PtrBool(false),
	})

	simplews.On("ping", func(c *ws.Customer, msg interface{}) {
		c.Emit("pong", msg)
	})

	simplews.OnConnect(func(consumer ws.Customer) error {
		consumer.Emit("hello_pipeline", "hello new customer !")
		return nil;
	})

	simplews.OnDisconnect(func(consumer ws.Customer, reason string) {
		fmt.Printf("disconnect: %s\n", reason)
	})

	http.Handle("/simple.ws", simplews)
	http.Handle("/", http.FileServer(http.Dir(".")))

	log.Println("listening on :8080")
	http.ListenAndServe("127.0.0.1:8080", nil)

}