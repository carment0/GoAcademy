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

func NewGame(p1 Player, p2 Player) *Game {
  return &Game{
    PlayerOne: p1,
    PlayerTwo: p2,
    CurrentPlayer: p1,
    Board: NewBoard(),
    TurnNum: 1,
  }
}

func (g *Game) printInfo() {
	fmt.Println("--------------------------------")
  fmt.Println("Turn #" + strconv.Itoa(g.TurnNum))
  fmt.Println(g.Board)
  fmt.Println("Current player:", g.CurrentPlayer.Name())
}

func (g *Game) Start() {
  fmt.Println("__Welcome to Connect Four in Go__")
  gameOver := false
  for gameOver == false {
    g.printInfo()
    if i, j, err := g.CurrentPlayer.GetMove(g.Board); err != nil {
      fmt.Println("Your input is invalid, please try again.")
    } else {
      g.Board.PlaceMark(i, j, g.CurrentPlayer.Mark())
      if g.Board.IsGameOver(i, j, g.CurrentPlayer.Mark()) {
        gameOver = true
      } else {
        g.switchPlayer()
        g.TurnNum += 1
      }
    }
  }
  fmt.Println(g.Board)
  fmt.Printf("Game Over %s won!", g.CurrentPlayer.Name())
}

func (g *Game) switchPlayer() {
  if g.CurrentPlayer == g.PlayerOne {
    g.CurrentPlayer = g.PlayerTwo
  } else {
    g.CurrentPlayer = g.PlayerOne
  }
}
