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
  Email string `json:email`
  Username string `json:username`
  Message string `json:message`
}

type Broker struct {
  // map of websocket connections
  ConnMap map[*websocket.Conn]bool
  Broadcast chan Payload
  // every websocket connect begins with a GET request will receive a request and server will immediately render a response. If we upgrade this request, we never return a response by keeping this connect alive. A long lived websocket connection!
  Upgrader *websocket.Upgrader
}

// create a controller, a high level fn which returns a fn
func (b *Broker) getHandleConnections() http.HandleFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    conn, err := b.Upgrader.Upgrade(w, r, nil)
    if err != nil {
      fmt.Println("Failed to upgrade GET to a websocket connection")
    }
    // defer = excute this fn when the handler fn is done
    defer conn.Close()

    // we want to store the connection
    b.ConnMap[conn] = true
    for {
      // constantly listen for Messages
      var p Payload
      if err := conn.ReadJSON(&p); err != nil {

      } else {
        fmt.Println("Incoming message:", p.Message)
        
      }
    }
  }
}
