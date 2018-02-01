package main

func main() {
	p1 := &HumanPlayer{"Momo", "X"}
	p2 := &HumanPlayer{"Donut", "O"}
	g := NewGame(p1, p2)
	g.Start()
}
