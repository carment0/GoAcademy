package main

import (
	"fmt"
	"strconv"
)

type Game struct {
  PlayerOne Player
  PlayerTwo Player
  CurrentPlayer Player
  Board *Board
  TurnNum int
}

// constructor
// return instance of Game
func NewGame(p1 Player, p2 Player) *Game {
  return &Game{
    PlayerOne: p1,
    PlayerTwo: p2,
    CurrentPlayer: p1,
    Board: NewBoard(),
    // need to add comma at the end
    TurnNum: 1,
  }
}

func (g *Game) printInfo() {
  // Using the package strconv, we can us .Itoa to return the result of ParseInt(s, 10, 0) converted to type int.
	fmt.Println("--------------------------------")
  fmt.Println("Turn #" + strconv.Itoa(g.TurnNum))
  fmt.Println(g.Board)
  fmt.Println("Current player:", g.CurrentPlayer.Name())
}

func (g *Game) Start() {
  fmt.Println("__Welcome to Tic Tac Toe in Go__")
  for !g.Board.IsOver() {
    g.printInfo()
    if i, j, err := g.CurrentPlayer.GetMove(g.Board); err != nil {
      fmt.Println("Your input is invalid, please try again.")
    } else {
      g.Board.PlaceMark(i, j, g.CurrentPlayer.Mark())
      g.switchPlayer()
      g.TurnNum += 1
    }
  }
  fmt.Println(g.Board)
  fmt.Println("Game Over")
}

func (g *Game) switchPlayer() {
  if g.CurrentPlayer == g.PlayerOne {
    g.CurrentPlayer = g.PlayerTwo
  } else {
    g.CurrentPlayer = g.PlayerOne
  }
}
