package main

import "fmt"

func main() {
	fmt.Print("Let's play Connect Four! \n")
	b := NewBoard()
	fmt.Println(b)
}