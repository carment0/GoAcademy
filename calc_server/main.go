package main

import (
  "fmt"
  "net/http"
)

func main() {
  // connect handler to server with .HandleFunc takes (route, handler)
  // .HandleFunc will give the two parameters for HandleAdd
  http.HandleFunc("/api/add", HandleAdd)
  // using the higher order fn
  http.HandleFunc("/api/subtract", OperationHandlerCreator(SUB))
	http.HandleFunc("/api/ ", OperationHandlerCreator(MUL))
	http.HandleFunc("/api/divide", OperationHandlerCreator(DIV))
  fmt.Println("Starting calculator server on port 3000")
  // listening for request and serve api endpoint
  err := http.ListenAndServe(":3000", nil)
  if err != nil {
    fmt.Println("Failed to start server", err)
  }
}

// To run:
// `go install`
// `calc server`
// On browser http://localhost:3000/api/add?lop=1&rop=1
