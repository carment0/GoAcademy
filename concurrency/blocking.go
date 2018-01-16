package main

import (
  "fmt"
  "time"
)

func BlockingFunc() {
  // forever for loop
  for {
    fmt.Println("I like to block you for a second.")
    time.Sleep(time.Second)
  }
}

func GetToWork() {
  for {
    fmt.Println("Please let me do my work")
    time.Sleep(time.Second)
  }
}

func main() {
  // without `go` GetToWork() will never be excuted bc BlockingFunc() doesn't end
  // `go` allows BlockingFunc to run on a different thread. when this file is run, two thread run simultaneously
  go BlockingFunc()
  GetToWork()
}

// Most laptops have 4 to 8 cores. Each core can run 8 threads
