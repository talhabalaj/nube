package main

import (
	"fmt"
	"nube/nube"
)

func main() {
	server := nube.CreateServer(4000)

	server.OnRequest(func(r *nube.Request, w *nube.ResponseWriter) {
		w.Json(map[string]any{
			"message": "Hello",
		})
	})

	fmt.Println("Starting listening on http://localhost:4000")
	server.Listen()
}