package parser

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
)

type Encoder struct {
	compression bool
	base64      bool
}

func NewEncoder(compression bool, base64 bool) *Encoder {
	return &Encoder{compression: compression, base64: base64}
}

func (e *Encoder) Encode(model * Model) (data []byte, err error) {

	if data, err = json.Marshal(*model); err != nil {
		return nil, err
	}

	if e.base64 {
		data = []byte(base64.StdEncoding.EncodeToString(data))
	}

	if e.compression {
		var b bytes.Buffer
		gz := gzip.NewWriter(&b)
		defer gz.Close()
		if _, err := gz.Write(data); err != nil {
			return nil, err
		}
		data = b.Bytes()
	}

	return data, nil
}