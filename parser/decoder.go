package parser

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
)

type Decoder struct {
	compression bool
	base64      bool
}

func NewDecoder(compression bool, base64 bool) *Decoder {
	return &Decoder{compression: compression, base64: base64}
}

func (d *Decoder) Decode(data []byte) (_ * Model, err error) {

	if d.compression {
		reader, err := gzip.NewReader(bytes.NewReader(data))
		if err != nil {
			return nil, err
		}
		defer reader.Close()
		data, err = ioutil.ReadAll(reader)
		if err != nil {
			return nil, err
		}
	}

	if d.base64 {
		data, err = base64.StdEncoding.DecodeString(string(data))
		if err != nil {
			return nil, err
		}
	}

	var model Model
	if err = json.Unmarshal(data, &model); err != nil {
		return nil, err
	}

	return &model, nil
}