package simplews_go

import (
	"github.com/antod3v/simplews_go/parser"
	"github.com/antod3v/simplews_go/ws"
)

type Server struct {
	opts    *Opts
	encoder * parser.Encoder
	decoder * parser.Decoder
	handles Handles
}

func New(opts *Opts) Server {
	
	return Server{
		opts: opts,
		encoder: parser.NewEncoder(*opts.Compression, *opts.Base64),
		decoder: parser.NewDecoder(*opts.Compression, *opts.Base64),
		handles: Handles{
			handleEvent: make(map[string][]func(customer * ws.Customer, msg interface{})),
		},
	}
}

func (h *Server) Everyone() {
	
}

func PtrInt(i int) *int {
	return &i
}

func PtrBool(i bool) *bool {
	return &i
}


