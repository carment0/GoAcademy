package main

import "fmt"

type Player interface {
	GetMove(b *Board) (int, int, error)
	Mark() string
	Name() string
}

type HumanPlayer struct {
	name string
	mark string
}

func (p *HumanPlayer) GetMove(b *Board) (int, int, error) {
	fmt.Print("Enter the column you want to drop your token to: ")
	var i int
	var n, err1 = fmt.Scanf("%d", &i)
	a, err2 := b.Drop(i)
	if err1 != nil || n != 1 {
		return 0, 0, err1
	} else if err2 != nil || n != 1 {
		return 0, 0, err2
	}

	fmt.Println("Your input:", a, i)
	return a, i, nil
}

func (p *HumanPlayer) Name() string {
	return p.name
}

func (p *HumanPlayer) Mark() string {
	return p.mark
}

type ComputerPlayer struct {
	name string
	mark string
}

func (cp *ComputerPlayer) GetMove(b *Board) (int, int, error) {
	move := cp.minimax(b, cp.Mark(), 1)
	i, j := move["i"], move["j"]
	return i, j, nil
}

func (cp *ComputerPlayer) Mark() string {
	return cp.mark
}

func (cp *ComputerPlayer) Name() string {
	return cp.name
}

func (cp *ComputerPlayer) minimax(b *Board, mark string, depth int) map[string]int {

}
