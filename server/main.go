package main

import (
	"fmt"
	"net/http"

	"github.com/dmls/discormb/server/handlers"
)

func main() {
	http.HandleFunc("/ws", handlers.WebSocketHandler)
	fmt.Println("Server is running on http://localhost:8080/ws")
	http.ListenAndServe(":8080", nil)
}
