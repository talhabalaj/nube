package nube

import (
	"fmt"
)

type Response struct {
	StatusCode int
	Headers    Header
	Version    string
	Body       []byte
}

func NewResponse() *Response {
	r := new(Response)
	r.StatusCode = 200
	r.Headers = *NewHeader()
	r.Headers.Set("server", "nube 0.0.1")
	return r
}

func (r *Response) WriteBody(body []byte, bodyType string) {
	r.Body = body
	r.Headers.Set("Content-Length", fmt.Sprint(len(body)))
	r.Headers.Set("Content-Type", bodyType)
}

func (r *Response) ToBytes() []byte {
	str := fmt.Sprintf("HTTP/1.1 %d %s\r\n", r.StatusCode, "OK")
	str += r.Headers.ToString()
	str += "\r\n"

	b := []byte(str)
	b = append(b, r.Body...)

	return b
}
