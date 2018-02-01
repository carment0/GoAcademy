package main

import (
  "fmt"
  "errors"
)

type Board struct {
  Grid [6][7]string
}

// init the board
func NewBoard() *Board {
  NewBoard := &Board{}

  for i := range NewBoard.Grid {
    for j := range NewBoard.Grid[i] {
      NewBoard.Grid[i][j] = "_"
    }
  }
  return NewBoard
}

// string method
func (b *Board) String() string {
  var boardStr string = "\n   0 1 2 3 4 5 6"
  row := 0
  for i := range b.Grid {
    boardStr += fmt.Sprintf("\n %d ", row)
    for j := range b.Grid[i] {
      boardStr += b.Grid[i][j] + " "
    }
    row += 1
  }
  return boardStr
}

// columns
func (b *Board) columns() [7][6]string {
  columns := [7][6]string{}
  for i := range b.Grid {
    for j := range b.Grid[i] {
      columns[j][i] = b.Grid[i][j]
    }
  }
  return columns
}


// token drop
func (b *Board) Drop(i int) (int, error) {
  columns := b.columns()
	for n := range columns[i] {
    if n == 0 && columns[i][n] != "_"{
      return 0, errors.New("no spots")
    } else if columns[i][n] != "_" {
			return n-1, nil
		} else if n == 5 {
      return n, nil
    }
	}
  return 0, errors.New("no spots")
}

// adding mark to board
func (b *Board) PlaceMark(i, j int, mark string) {
  b.Grid[i][j] = mark
}

// checking if over
func (b *Board) IsGameOver(i, j int, mark string) bool {
  if b.isTied() || b.isWon(i, j, mark) {
    return true
  }
  return false
}

func (b Board) Copy() *Board {
  return &b
}

func (b *Board) EmptySpaces() [][2]int {
  emptySpaces := [][2]int{}
  for i := range b.Grid {
    for j := range b.Grid {
      if b.Grid[i][j] == "_" {
        emptySpaces = append(emptySpaces, [2]int{i, j})
      }
    }
  }
  return emptySpaces
}

// checking if tied
func (b *Board) isTied() bool {
  for i := range b.Grid {
    for j:= range b.Grid {
      if b.Grid[i][j] == "_" {
        return false
      }
    }
  }
  return true
}

// checking if player won
func (b *Board) isWon(i, j int, mark string) bool {
  x := b.rowCount(i, j, mark)
  y := b.columnCount(i, j, mark)
  l := b.leftDiagonalCount(i, j, mark)
  r := b.rightDiagonalCount(i, j, mark)
  directionCount := [4]int{x, y, l, r}
  for n := range directionCount {
    if directionCount[n] == 5 {
      return true
    }
  }
  return false
}

func (b *Board) countAlongDirection(i, j int, mark string, count, xDelta, yDelta int) int {
  if i < 0 || j < 0 || i > 5 || j > 6 {
    return count
  }
  if b.mark(i, j) != mark {
    return count
  }
  count += 1
  return b.countAlongDirection(i+xDelta, j+yDelta, mark, count, xDelta, yDelta)
}

func (b *Board) rowCount(i, j int, mark string) int {
  left := b.countAlongDirection(i, j, mark, 0, 0, -1)
  right := b.countAlongDirection(i, j, mark, 0, 0, 1)
  return left + right
}

func (b *Board) columnCount(i, j int, mark string) int {
  up := b.countAlongDirection(i, j, mark, 0, -1, 0)
  down := b.countAlongDirection(i, j, mark, 0, 1, 0)
  return up + down
}

func (b *Board) leftDiagonalCount(i, j int, mark string) int {
  upLeft := b.countAlongDirection(i, j, mark, 0, -1, -1)
  downRight := b.countAlongDirection(i, j, mark, 0, 1, 1)
  return upLeft + downRight
}

func (b *Board) rightDiagonalCount(i, j int, mark string) int  {
  upRight := b.countAlongDirection(i, j, mark, 0, 1, -1)
  downLeft := b.countAlongDirection(i, j, mark, 0, -1, 1)
  return upRight + downLeft
}

func (b *Board) mark(i, j int) string {
  return b.Grid[i][j]
}
