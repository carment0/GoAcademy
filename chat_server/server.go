package main

import (
	"fmt"
	"net/http"
  "github.com/gorilla/websocket"
)

func main() {
	// FileServer will serve all the files in the public folder as a file server
	fs := http.FileServer(http.Dir("public"))

	b := &Broker{
		ConnMap:   make(map[*websocket.Conn]bool),
		Broadcast: make(chan Payload),
		Upgrader:  &websocket.Upgrader{},
	}

  // this is a blocking call (forever loop), so we need to add a new thread!
  go b.handleBroadcast()

	http.Handle("/", fs)
	// HandleFunc will call when the user hit the endpoint, there will be a new HandleFunc in the getHandleConnections fn.
	// But there is only one broker. Which means only one Broadcast channel.
	http.HandleFunc("/ws", b.getHandleConnections())

	fmt.Println("Starting server on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
