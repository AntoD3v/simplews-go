package simplews_go

import "net/http"

type Opts struct {
	ReadBufferSize  *int  `default:"1024"`
	WriteBufferSize *int  `default:"1024"`
	Compression     *bool `default:"true"`
	Base64          *bool `default:"false"`
	CheckOrigin     *func(r *http.Request) bool
}
