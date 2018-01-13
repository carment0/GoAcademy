package main

// Go’s structs are typed collections of fields. They’re useful for grouping data together to form records. Object like???
type Board struct {
  // We are creating a array called Grid: [n]T
  // n = number of elements to storage
  // T = type of elements
  // the string is either "x", "o"
  // Grid is the key
  // array is the value
  Grid [3][3]string
}

// Let's create a constructor function
// the return value is a pointer to a Board struct instance
func NewBoard() *Board {
  NewBoard := &Board{}
  // A "for" statement with a "range" clause iterates through all entries of an array, slice, string or map, or values received on a channel.
  for i := range NewBoard.Grid  {
    for j := range NewBoard.Grid[i] {
      NewBoard.Grid[i][j] = "_"
    }
  }
  return NewBoard
}

func (b *Board) PlaceMark(i, j int, mark string) {
  b.Grid[i][j] = mark
}

func (b *Board) IsOver() bool {
  return b.isWon() || b.isTied()
}

// to string method
func (b *Board) String() string {
  var boardStr string = "\n "
  for i := range b.Grid {
    for j := range b.Grid {
      boardStr += b.Grid[i][j] + " "
    }
    boardStr += "\n "
  }
  return boardStr
}

// Capitalize funcation name = Public
// Lower-case function name = Private
func (b *Board) isTied() bool {
  for i := range b.Grid {
    for j := range b.Grid[i] {
      if b.Grid[i][j] == "_" {
        return false
      }
    }
  }
  return true
}

func (b *Board) isWon() bool {
  if b.Winner() == "" {
    return false
  }
  return true
}

func (b *Board) Winner() string {
  xStreak := [3]string {"X", "X", "X"}
  oStreak := [3]string {"O", "O", "O"}

  rows := b.rows()
  for i := range rows {
    if rows[i] == xStreak {
      return "X"
    }

    if rows[i] == oStreak {
      return "O"
    }
  }

  columns := b.columns()
  for i := range columns {
    if columns[i] == xStreak {
      return "X"
    }

    if columns[i] == oStreak {
      return "O"
    }
  }

  diagonals := b.diagonals()
  for i := range diagonals {
    if diagonals[i] == xStreak {
      return "X"
    }

    if diagonals[i] == oStreak {
      return "O"
    }
  }
  return ""
}

func (b *Board) rows() [3][3]string {
  return b.Grid
}

func (b *Board) columns() [3][3]string {
  columns := [3][3]string{}
  for i := range b.Grid {
    for j := range b.Grid[i] {
      columns[j][i] = b.Grid[i][j]
    }
  }
  return columns
}

func (b *Board) diagonals() [2][3]string {
  diagonals := [2][3]string{}
  for i := range b.Grid {
    diagonals[0][i] = b.Grid[i][i]
  }

  for i := range b.Grid {
    diagonals[1][i] = b.Grid[i][2-i]
  }
  return diagonals
}



// _______________________Bonus______________________________
func (b *Board) emptySpaces() [][2]int {
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

func (b Board) Copy() *Board {
	return &b
}
