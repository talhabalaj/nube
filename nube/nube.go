package nube

import (
	"fmt"
	"net"
)

type RequestHandler func(*Request, *ResponseWriter)

type Nube struct {
	Port           int
	listener       net.Listener
	requestHandler RequestHandler
	errorHandler   func(error, *Request, *ResponseWriter)
}


func defaultErrorHandle(err error, req *Request, writer *ResponseWriter) {

}

func CreateServer(port int) *Nube {
	n := new(Nube)
	n.errorHandler = defaultErrorHandle
	n.Port = port

	return n
}

func (n *Nube) OnRequest(r RequestHandler) {
	n.requestHandler = r
}

func (n *Nube) handleConnection(conn net.Conn) {
	req, err := NewRequest(conn)
	resWriter := NewResponseWriter(conn)

	if err != nil {
		n.errorHandler(err, req, resWriter)
	} else {	
		n.requestHandler(req, resWriter)
	}
}

func (n *Nube) Listen() {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", n.Port))

	if err == nil {
		n.listener = ln

		for {
			conn, err := n.listener.Accept()

			if err == nil {
				go n.handleConnection(conn)
			}
		}
	} else {
		panic(err)
	}
}
