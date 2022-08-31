# SimpleWS

SimpleWS inspired by socket.io is a websocket librairy written in Golang. The project was created because today, it's not exist a socket.io v4 librairy in Golang

**_Warning_**: SimpleWS is'nt work with socket.io, the protocol is completely different !!!

## Create an instance

```go
    simplews := simplews_go.New(&simplews_go.Opts{
  ReadBufferSize:  simplews_go.PtrInt(1024),
  WriteBufferSize: simplews_go.PtrInt(1024),
  Compression:     simplews_go.PtrBool(false),
  Base64:          simplews_go.PtrBool(true),
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
```

## Websocket Protocol

SimpleWS is ... simple. Each message is encoded in JSON with this format :

| Key     | Type of value |                         Description                         |
| :------ | :-----------: | :---------------------------------------------------------: |
| Type    | integer [0-3] |      connect: 0, disconnect: 1, event: 2 and error: 3       |
| Event   |    string     |            name of your event ("ping" or "pong")            |
| Data    |    object     |                      related to event                       |
| Message |    string     | Only for the error message or the reason for disconnection. |

**Note**: All of useless keys can be delete

Example of websocket message with SimpleWS

```json
{
  "type": 2,
  "event": "new_user",
  "data": {
    "name": "Antoine",
    "age": 20
  }
}
```

Finally, the JSON message will be encoded with base64 and compressed with GZIP.
