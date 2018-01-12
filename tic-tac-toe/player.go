package main

import "fmt"

// we are using interface bc we will have different type of players
// we want it to be generic discription for a list of struct
// any struct that has a the same functions listed under the Player interfere is a type of Player
type Player interface {
  // what are the fn all you player can do?
  GetMove(b *Board) (int, int, error)
  Mark() string
  Name() string
}

// Struct is like a js class
// Once you define a struct your cannot add new keys
// if you print a instance of the type you get a object
type HumanPlayer struct {
  name string
  mark string
}

// invoked by p.GetMove(b)
// p = receiver
// b = arugments
// (int, int, error) = expected outcome of this function
func (p *HumanPlayer) GetMove(b *Board) (int, int, error) {
  fmt.Print("Enter a position: ")
  var i, j int
  // the condition of the if statement is determined by the return value of .Scanf
  // .Scanf asks for user input (will be shipped to mailing address)
  // .Scanf arguments are the point receiver of input i and j (mailing address)
  // %d is a verb for base 10 integers, "%d %d" requird format the user needs to type
  // .Scanf analyze the user input and returns:
  // n = number for char input, err = error due to the inputs
  // you can also get user input and setting outside the if statement. The ; is required when you want to do that on the same line
  if n, err := fmt.Scanf("%d %d", &i, &j); err != nil || n !=2 {
    return 0, 0, err
  }

  fmt.Println("Your input:", i, j)
  return i, j, nil
}

// Getters
func (p *HumanPlayer) Mark() string {
  return p.mark
}

func (p *HumanPlayer) Name() string {
  return p.name
}

// Bonus
type ComputerPlayer struct {
  name string
  mark string
}
