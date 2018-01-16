package main

import (
  "fmt"
  "time"
)

type Dog struct {
  Name string
  Age int
}

// takes two channels in and out
func AgeTheDog(in chan *Dog, out chan *Dog) {
  // this `for` is a forever loop that listens to the in channel
  // when something is recieved from the in channel, it will be stored in the dog variable
  for dog := range in {
    dog.Age += 1
    time.Sleep(time.Second)
    // send out the dog
    out <- dog
  }
}

// main fn is the primary thread
func main() {
  // create channels
  inTunnel := make(chan *Dog)
  outTunnel := make(chan *Dog)

  // the second thread
  go AgeTheDog(inTunnel, outTunnel)

  // make a dog
  loki := &Dog{
    Name: "Loki",
    Age: 0,
  }

  for {
    // send Loki to the tunnel for aging!
    inTunnel <- loki
    <- outTunnel

    fmt.Println(loki)
  }
}
