package main

import "fmt"

// we are using interface bc we will have different type of players
// we want it to be generic discription for a list of struct
// any struct that has a the same functions listed under the Player interface is a type of Player
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

// _______________________Bonus______________________________
type ComputerPlayer struct {
  name string
  mark string
}

func (cp *ComputerPlayer) GetMove(b *Board) (int, int, error) {
	move := cp.minimax(b, cp.Mark(), 1)
	i, j := move["i"], move["j"]
	return i, j, nil
}

// Getters
func (cp *ComputerPlayer) Mark() string {
	return cp.mark
}

func (cp *ComputerPlayer) Name() string {
	return cp.name
}

// we want the return value of the recursion to be a hash map e.g. {pos1: #, pos2: #, score: #}
func (cp *ComputerPlayer) minimax(b *Board, mark string, depth int) map[string]int {
  // Basecase is when the game is over, you will just pass the score not the position bc you won!
	if b.IsOver() {
		// make() is the only way to create a hash map, must give it type(s). String as the key and int as the value
		score := make(map[string]int)
    // check if the ComputerPlayer is the winner and return final score
		if b.Winner() == cp.Mark() {
			score["value"] = 10 - depth
		} else {
			score["value"] = depth - 10
		}
		return score
	}

  // create an array to store all of hash maps
	scores := []map[string]int{}
  // when doing a range on the nested array of avaiable positions, each output is a index, element. In this case we only need the element (position).
	for _, pos := range b.emptySpaces() {
    // Make a copy of the current board to place a potential mark. No need to define type here when creating a variable bc we are using `:=` !!!!
		newBoard := b.Copy()
		i, j := pos[0], pos[1]
		newBoard.PlaceMark(i, j, mark)

    // Why not use make()? bc we are creating a variable not making a hash map. Why we need to create another hash map? bc we will be recieving a hash map from the if statement above due to recursion
		var score map[string]int
    // switch marks to demostrate each possiblities
		if mark == "X" {
			score = cp.minimax(newBoard, "O", depth+1)
		} else {
			score = cp.minimax(newBoard, "X", depth+1)
		}

    // once you are done with the recursion, you set the position
		score["i"] = i
		score["j"] = j
    // score e.g. {i: 1, j: 2, score: -12} the positions and the final score due to that move
		scores = append(scores, score)
	}

  // we want the highest score
	if mark == cp.Mark() { // max
		maxScore := scores[0]
		for _, s := range scores {
			if maxScore["value"] < s["value"] {
				maxScore = s
			}
		}
		return maxScore
    // we want the lowests
	} else { // min
		minScore := scores[0]
		for _, s := range scores {
			if minScore["value"] > s["value"] {
				minScore = s
			}
		}
		return minScore
	}
}
