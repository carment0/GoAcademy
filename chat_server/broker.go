package main

import (
	"fmt"
	"net/http"
	// how to import external libraries
	// you need to install `go dep`. dep is a dependency management for Golang (similar to npm and bundle). This will pull a source code from an open source repo!
	// in the terminal:
	// `dep init` - will check your source code to see if there is any missing libraries and it will create a file (go package) that will list all the things you need
	// `dep ensure` - install all the missing dependencies from the package into a folder called vendor
	"github.com/gorilla/websocket"
)

type Payload struct {
	// here we will add a json tag.
	// when your server receive a json from the client, you want to unload the json into a struct. you just add a json tag to get the value
	// e.g. received this json from client {email: "momo@to.com", username: "momoto", message: "meow"}
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

type Broker struct {
	// map of websocket connections
	ConnMap   map[*websocket.Conn]bool
	Broadcast chan Payload
	// every websocket connect begins with a GET request will receive a request and server will immediately render a response. If we upgrade this request, we never return a response by keeping this connect alive. A long lived websocket connection!
	Upgrader *websocket.Upgrader
}

// take payload from the Broadcast channel and Broadcast it to everyone who is connected to the server
func (b *Broker) handleBroadcast() {
	// listen to the Broadcast forever
	for payload := range b.Broadcast {
		// iterate through the connections
		for conn := range b.ConnMap {
			err := conn.WriteJSON(payload)
			if err != nil {
				fmt.Println("Failed to write JSON to websocket connection.")
				conn.Close()
				delete(b.ConnMap, conn)
			}
		}
	}
}

// create a controller, a high level fn which returns a fn
func (b *Broker) getHandleConnections() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    // lets check if we can upgrade the GET!
		conn, err := b.Upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("Failed to upgrade GET to a websocket connection")
		}
		// defer = excute this fn when the handler fn is done
		defer conn.Close()

		// we want to store the connection if there was no GET error
		b.ConnMap[conn] = true
    // lets check if the payload is correct!
		for {
			// constantly listen for Messages
			var p Payload
			if err := conn.ReadJSON(&p); err != nil {
				if websocket.IsUnexpectedCloseError(err) {
					fmt.Println("Client connection has closed")
				} else {
					fmt.Println("Encounter websocket error")
				}
				// delete live connention
				delete(b.ConnMap, conn)
				return
			} else {
				fmt.Println("Incoming message:", p.Message)
				// send the message to the channel Broadcast
				// one channel but multiple connection!
				b.Broadcast <- p
			}
		}
	}
}

Questions?
what does ConnMap, Broadcast look like?
line 37 what does conn.WriteJSON(payload) mean?
line 40 does closing the connection end the go thread?
line 50 what does b.Upgrader.Upgrade(w, r, nil) mean?
line 55 what does defer conn.Close() mean?
line 62 what does conn.ReadJSON(&p) mean? IsUnexpectedCloseError?
b.Broadcast <- p?
