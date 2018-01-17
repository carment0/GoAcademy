package main

import (
  "fmt"
  "net/http"
)

func main() {
  // FileServer will serve all the files in the public folder as a file server
  fs := http.FileServer(http.Dir("public"))
  http.Handle("/", fs)

  fmt.Println("Starting server on port 8000")
  err := http.ListenAndServe(":8000", nil)
  if err != nil {
    fmt.Println("ListenAndServe: ", err)
  }
}
