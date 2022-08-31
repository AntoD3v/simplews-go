package parser

type Type int

const (
	Connected  Type = 0
	Disconnect Type = 1
	Event      Type = 2
	Error      Type = 3
)

// {type: 1, "event":"message","data":{"message":"Hello World"}}
type Model struct {
	Type    Type         `json:"type"`
	Event   *string      `json:"event"`
	Data    *interface{} `json:"data"`
	Message *string      `json:"message"` // error message or disconnect reason
}