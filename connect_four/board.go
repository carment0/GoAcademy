package main

import "fmt"

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
  var boardStr string = "\n   1 2 3 4 5 6 7"
  row := 1
  for i := range b.Grid {
    boardStr += fmt.Sprintf("\n %d ", row)
    for j := range b.Grid[i] {
      boardStr += b.Grid[i][j] + " "
    }
    row += 1
  }
  return boardStr
}
//
//// adding mark to board
//func (b *Board) PlaceMark(i, j int, mark string) {
//  b.Grid[i][j] = mark
//}
//
//// checking if tied
//func (b *Board) isTied() bool {
//  for i := range b.Grid {
//    for j:= range b.Grid {
//      if b.Grid[i][j] == "_" {
//        return false
//      }
//    }
//  }
//  return true
//}
//
//// checking if over
//func (b *Board) isOver() bool {
//  if b.isTied() || b.isWon() {
//    return true
//  }
//  return false
//}
//
//// checking if anyone won
//func (b *Board) isWon() bool {
//  if b.Winner() == "" {
//    return false
//  }
//  return true
//}
//
//// get array of row positions
//func (b *Board) rows() [5][7]string {
//  return b.Grid
//}
//
//// get array of column positions
//func (b *Board) columns() [7][5]string {
//  columns := [7][5]string{}
//  for i := range b.Grid {
//    for j := range b.Grid {
//      columns[j][i] = b.Grid[i][j]
//    }
//  }
//  return columns
//}
//
//// check if there is a diagonal win
//func (b *Board) diagonalBottomRightWin(i, j int, mark string) bool {
//  var count = 0
//  for n := 0; n < 4; i++ {
//    if i+n > 5 || j+n > 7 {
//      return false
//    }
//    if b.Grid[i+n][j+n] == mark {
//      count += 1
//    } else {
//      return false
//    }
//  }
//  return true
//}
//
//func (b *Board) diagonalBottomLeftWin(i, j int, mark string) bool {
//  var count = 0
//  for n := 0; n < 4; i++ {
//    if i+n > 5 || j+n < 0 {
//      return false
//    }
//    if b.Grid[i+n][j-n] == mark {
//      count += 1
//    } else {
//      return false
//    }
//  }
//  return true
//}
//
//func (b *Board) diagonalTopRightWin(i, j int, mark string) bool {
//  var count = 0
//  for n := 0; n < 4; i++ {
//    if i-n < 0 || j+n > 7 {
//      return false
//    }
//    if b.Grid[i-n][j+n] == mark {
//      count += 1
//    } else {
//      return false
//    }
//  }
//  return true
//}
//
//func (b *Board) diagonalTopLeftWin(i, j int, mark string) bool {
//  var count = 0
//  for n := 0; n < 4; i++ {
//    if i+n < 0 || j+n < 0 {
//      return false
//    }
//    if b.Grid[i-n][j-n] == mark {
//      count += 1
//    } else {
//      return false
//    }
//  }
//  return true
//}
//
//func (b *Board) countAlongDirection(i, j int, mark string, count int) int {
//
//}