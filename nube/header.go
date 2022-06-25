package nube

import (
	"fmt"
	"strings"
)

const HEADER_SEP = ":"

type Header struct {
	headers map[string]string
}

func (header *Header) Get(name string) string {
	return header.headers[strings.ToLower(name)]
}

func (header *Header) Set(name string, value string) {
	header.headers[strings.ToLower(name)] = value
}

func NewHeader() *Header {
	r := new(Header)
	r.headers = make(map[string]string)

	return r
}

func (header *Header) ToString() string {
	str := ""
	
	for key, value := range header.headers {
		str += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	return str
}