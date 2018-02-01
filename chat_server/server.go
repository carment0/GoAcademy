package main

import (
	"fmt"
	"net/http"
  "github.com/gorilla/websocket"
)

func main() {
	// FileServer will serve all the files in the public folder as a file server
	fs := http.FileServer(http.Dir("public"))

  // create a broker
	b := &Broker{
		ConnMap:   make(map[*websocket.Conn]bool),
		Broadcast: make(chan Payload),
    // Upgrade upgrades the HTTP server connection to the WebSocket protocol.
		Upgrader:  &websocket.Upgrader{},
	}

  // this is a blocking call (forever loop), so we need to add a new thread!
  go b.handleBroadcast()

  // Think of two functions, one is `FeedAnimal` and another one is `FeedCat`... `FeedAnimal` expects inputs to be ANYTHING that qualifies to be an animal. `FeedCat` expects inputs to be CAT ONLY CAT.
	// `Handle()` is equivalent to `FeedAnimal`
	// `HandleFunc()` is equivalent to `FeedCat`
	// The reason why `Handle` was used, because `FileServer` is something that qualifies to be a `Handler`
	// It's like when do you use FeedAnimal vs FeedCat... Let's say I am only given two choices... and I am trying to feed Loki
	// Obviously FeedCat won't work
	// The only option I have left is FeedAnimal
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


Questions?
line 18 explain websocket.Upgrader{}
why need two threads?
