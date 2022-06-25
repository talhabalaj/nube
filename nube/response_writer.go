package nube

import (
	"encoding/json"
	"net"
)

type ResponseWriter struct {
	conn net.Conn
	response* Response
}

func NewResponseWriter(conn net.Conn) *ResponseWriter {
	r := new(ResponseWriter)
	r.conn = conn
	r.response = NewResponse()

	return r
}

func (writer *ResponseWriter) SetHeader(key string, value string) {
	writer.response.Headers.Set(key, value)
}

func (writer *ResponseWriter) GetHeader(key string) string {
	return writer.response.Headers.Get(key)
}

func (writer *ResponseWriter) Status(code int) *ResponseWriter {
	writer.response.StatusCode = code
	return writer
}


func (writer* ResponseWriter) Send(str string) {
	writer.response.WriteBody([]byte(str), "text/plain")
	writer.conn.Write(writer.response.ToBytes())
}

func (writer* ResponseWriter) Json(a any) {
	parsed, err := json.Marshal(a)

	if err != nil {
		panic(err)
	}

	writer.response.WriteBody(parsed, "application/json")
	writer.conn.Write(writer.response.ToBytes())
}
