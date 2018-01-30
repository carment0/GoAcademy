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
	fmt.Print("Enter a position: ")
	var i, j int
	var n, err = fmt.Scanf("%d %d", &i, &j)
	if err != nil || n != 2 {
		return 0, 0, err
	}

	fmt.Println("Your input:", i, j)
	return i, j, nil
}

func (p *HumanPlayer) Name() string {
	return p.name
}

func (p *HumanPlayer) Mark() string {
	return p.mark
}


