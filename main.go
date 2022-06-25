package main

import (
	"fmt"
	"nube/nube"
)

func main() {
	server := nube.CreateServer(4000)

	server.OnRequest(func(r *nube.Request, w *nube.ResponseWriter) {
		w.Send(fmt.Sprintf("This is %s @ %s", r.Method, r.Path))
	})

	fmt.Println("Starting listening on http://localhost:4000")
	server.Listen()
}